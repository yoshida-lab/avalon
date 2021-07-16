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

type (
	SSHExecutor struct {
		UseSCP bool
		SSHConfig
		localExecutor
	}
)

func (sshExec *SSHExecutor) SyncFileRemote(srcDest map[string]string, currentDir string, fromLocal bool, opts *CopyOptions) (err error) {
	panic("implement me")
}

func (sshExec *SSHExecutor) ExecRemote(currentDir string, commands *[]string) (stdOut <-chan string, stdErr <-chan string, err <-chan error) {
	panic("")
}
