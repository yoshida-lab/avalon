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

package scp

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path/filepath"
	"time"
)

// Upload uploads sourceFile to remote machine like native scp console app.
func Upload(localFilePath, remoteDir string, permissions os.FileMode, session *ssh.Session, timeout ...time.Duration) error {
	defer func(session *ssh.Session) {
		_ = session.Close()
	}(session)

	fileName := filepath.Base(localFilePath)
	destFilePath := filepath.Join(remoteDir, fileName)

	src, srcErr := os.Open(localFilePath)
	if srcErr != nil {
		return srcErr
	}

	srcStat, statErr := src.Stat()

	if statErr != nil {
		return statErr
	}

	w, err := session.StdinPipe()
	if err != nil {
		return err
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	copyF := func() error {
		_, err := fmt.Fprintf(w, "C%#o %d %s\n", permissions, srcStat.Size(), fileName)
		if err != nil {
			return err
		}

		if err = checkResponse(stdout); err != nil {
			return err
		}

		if srcStat.Size() > 0 {
			_, err = io.Copy(w, src)
			if err != nil {
				return err
			}
		}

		_, err = fmt.Fprint(w, "\x00")
		if err != nil {
			return err
		}

		if err = checkResponse(stdout); err != nil {
			return err
		}

		return nil
	}

	copyErrC := make(chan error, 1)
	go func() {
		defer func(w io.WriteCloser) {
			_ = w.Close()
		}(w)
		copyErrC <- copyF()
	}()

	runErr := session.Run(fmt.Sprintf("scp -tr %s", destFilePath))

	if len(timeout) > 0 {
		select {
		case err = <-copyErrC:
		case <-time.After(timeout[0]):
			err = errors.New(fmt.Sprintf("file transfer timeout (%v)", timeout[0])) // timed out
		}
	} else {
		err = <-copyErrC
	}

	if runErr != nil {
		return errors.New(fmt.Sprintf("upload file failed! 'execution error': %+v; 'transfer error': %+v", runErr, err)) // completed normally
	}
	return err
}

// Download Copy a file from the remote to the given writer. The passThru parameter can be used
// to keep track of progress and how many bytes that were download from the remote.
// `passThru` can be set to nil to disable this behaviour.
func Download(remoteFilePath, localDir string, permissions os.FileMode, session *ssh.Session, timeout ...time.Duration) error {
	defer func(session *ssh.Session) {
		_ = session.Close()
	}(session)

	fileName := filepath.Base(remoteFilePath)
	localFilePath := filepath.Join(localDir, fileName)

	w, localErr := os.OpenFile(localFilePath, os.O_WRONLY|os.O_CREATE, permissions)
	if localErr != nil {
		return localErr
	}

	copyF := func() error {
		r, err := session.StdoutPipe()
		if err != nil {
			return err
		}

		in, err := session.StdinPipe()
		if err != nil {
			return err
		}
		defer func(in io.WriteCloser) {
			_ = in.Close()
		}(in)

		err = session.Start(fmt.Sprintf("scp -f %q", remoteFilePath))
		if err != nil {
			return err
		}

		err = Ack(in)
		if err != nil {
			return err
		}

		res, err := ParseResponse(r)
		if err != nil {
			return err
		}

		infos, err := res.ParseFileInfos()
		if err != nil {
			return err
		}

		err = Ack(in)
		if err != nil {
			return err
		}

		_, err = CopyN(w, r, infos.Size)
		if err != nil {
			return err
		}

		err = Ack(in)
		if err != nil {
			return err
		}

		err = session.Wait()
		if err != nil {
			return err
		}
		return nil
	}

	copyErrC := make(chan error, 1)
	go func() {
		defer func(w io.WriteCloser) {
			_ = w.Close()
		}(w)
		copyErrC <- copyF()
	}()

	if len(timeout) > 0 {
		select {
		case err := <-copyErrC:
			return err // completed normally
		case <-time.After(timeout[0]):
			return errors.New(fmt.Sprintf("file transfer timeout (%v)", timeout[0])) // timed out
		}
	} else {
		err := <-copyErrC
		return err
	}
}
