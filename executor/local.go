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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type LocalExecutor struct{}

func (l *LocalExecutor) SyncFileLocal(src, dest string, opts *CopyOptions) error {
	panic("implement me")
}

func (l *LocalExecutor) ExecLocal(currentDir string, commands *[]string, fn func(error)) {
	var fileHandle []*os.File
	// close out files
	defer func() {
		for _, f := range fileHandle {
			_ = f.Close()
		}
	}()

	logPath := filepath.Join(currentDir, "logs")
	// make logs dir if not exist with perm set to 0777
	// the umask will take away the unwanted permissions,
	// leaving you with the right permissions.
	if err := os.MkdirAll(logPath, 0777); err != nil {
		fn(&Error{IOError, err})
		return
	}

	nowStr := time.Now().Format("2006-01-02-15:04:05")
	for i, cmd := range *commands {
		count := i + 1
		stdoutFile := filepath.Join(logPath, fmt.Sprintf("%v.%v.out", nowStr, count))
		stderrFile := filepath.Join(logPath, fmt.Sprintf("%v.%v.err", nowStr, count))

		outFile, err := os.Create(stdoutFile)
		fileHandle = append(fileHandle, outFile)
		if err != nil {
			fn(&Error{IOError, err})
			return
		}

		errFile, err := os.Create(stderrFile)
		fileHandle = append(fileHandle, errFile)
		if err != nil {
			fn(&Error{IOError, err})
			return
		}

		cmdAndArgs := strings.Fields(cmd)
		_exec := exec.Command(cmdAndArgs[0], cmdAndArgs[1:]...)
		_exec.Dir = currentDir
		_exec.Stdout = outFile
		_exec.Stderr = errFile

		// exec command
		if err := _exec.Start(); err != nil {
			fn(&Error{CmdExecutionError, err})
			return
		}

		// wait until done
		if err := _exec.Wait(); err != nil {
			fn(&Error{CmdExecutionError, err})
			return
		}
	}
}
