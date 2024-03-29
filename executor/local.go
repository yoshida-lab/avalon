/*
 * Copyright (c) 2021 TsumiNa
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package executor

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const msgBufferSize = 1024

type localExecutor struct{}

func NewLocalExecutor() *localExecutor {
	return &localExecutor{}
}

func (l *localExecutor) SyncFileRemote(map[string]string, string, bool, *CopyOptions) (err error) {
	panic("Local executor cannot be used to transform files to remote")
}

func (l *localExecutor) ExecRemote(currentDir string, commands *[]string) (stdOut <-chan string, stdErr <-chan string, err <-chan error) {
	panic("Local executor cannot be used to execute commands om remote")
}

func copyOrMove(src, dest string, perm fs.FileMode, move bool) (err error) {
	if move {
		if err = os.Rename(src, dest); err != nil {
			return
		}
		return os.Chmod(dest, perm)

	}

	sf, err := os.Open(src)
	if err != nil {
		return
	}
	defer func(sf *os.File) {
		err = sf.Close()
	}(sf)

	df, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		return
	}
	defer func(df *os.File) {
		err = df.Close()
	}(df)

	_, err = io.Copy(df, sf)
	return
}

func syncOne(src, dest string, opts *CopyOptions) (err error) {
	// convert paths to shortest
	src, err = filepath.Abs(src)
	if err != nil {
		return
	}
	dest, err = filepath.Abs(dest)
	if err != nil {
		return
	}

	// create dest dirs
	// this will create parents if not exist
	if err = os.MkdirAll(dest, 0755); err != nil {
		return
	}

	// check src file then apply copy or move base on the "opts"
	info, err := os.Stat(src)
	if os.IsNotExist(err) {
		return
	}
	if info.IsDir() {
		err = filepath.WalkDir(src, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(src, path)
			if info.IsDir() {
				if err != nil {
					return err
				}
				if err = os.MkdirAll(filepath.Join(dest, rel), 0755); err != nil {
					return err
				}
			}
			return copyOrMove(path, filepath.Join(dest, rel), opts.FileMode, opts.MoveContent)
		})
	} else {
		err = copyOrMove(src, filepath.Join(dest, filepath.Base(src)), opts.FileMode, opts.MoveContent)
	}
	return
}

func (l *localExecutor) SyncFileLocal(srcDest map[string]string, currentDir string, opts *CopyOptions) (err error) {
	err = os.Chdir(currentDir)
	if err != nil {
		return
	}
	for src, dest := range srcDest {
		err = &Error{IOError, syncOne(src, dest, opts)}
	}
	return
}

func (l *localExecutor) ExecLocal(currentDir string, commands *[]string) (<-chan string, <-chan string, <-chan error) {
	stdOutChan := make(chan string)
	stdErrChan := make(chan string)
	errChan := make(chan error)

	go func() {
		defer close(stdErrChan)
		defer close(stdOutChan)
		defer close(errChan)

		for _, cmd := range *commands {

			cmdAndArgs := strings.Fields(cmd)
			_exec := exec.Command(cmdAndArgs[0], cmdAndArgs[1:]...)
			_exec.Dir = currentDir

			stdout, err := _exec.StdoutPipe()
			if err != nil {
				errChan <- &Error{CmdExecutionError, err}
			}
			stderr, err := _exec.StderrPipe()
			if err != nil {
				errChan <- &Error{CmdExecutionError, err}
			}

			// exec command
			if err = _exec.Start(); err != nil {
				errChan <- &Error{CmdExecutionError, err}
			}

			var resWg sync.WaitGroup
			resWg.Add(2)

			nowStr := time.Now().Format("2006-01-02-15:04:05")
			stdOutChan <- fmt.Sprintf("%v (%v):", cmd, nowStr)
			go func() {
				stdoutScanner := bufio.NewScanner(stdout)
				for stdoutScanner.Scan() {
					stdOutChan <- stdoutScanner.Text()
				}
				stdOutChan <- " "
				resWg.Done()
			}()

			stdErrChan <- fmt.Sprintf("%v (%v):", cmd, nowStr)
			go func() {
				stderrScanner := bufio.NewScanner(stderr)
				for stderrScanner.Scan() {
					stdErrChan <- stderrScanner.Text()
				}
				stdErrChan <- " "
				resWg.Done()
			}()

			// wait until done
			if err = _exec.Wait(); err != nil {
				errChan <- &Error{CmdExecutionError, err}
			}
			resWg.Wait()
		}
	}()

	return stdOutChan, stdErrChan, errChan
}
