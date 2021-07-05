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
	"testing"
)

func TestLocalExecutor_ExecLocal(t *testing.T) {
	type args struct {
		currentDir string
		commands   *[]string
		fn         func(error)
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"make log dir error",
			args{
				"/not_exist_dir",
				&[]string{"ls", "-l"},
				func(err error) {
					if err != nil {
						panic(err)
					}
				},
			},
			"IOError (0): mkdir /not_exist_dir: read-only file system",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				} else {
					if e, ok := r.(*Error); ok {
						if e.String() != tt.want {
							t.Errorf("panic with unexpected information: %v", r)
						}
					} else {
						t.Errorf("expecte type '*executor.Error' got: %T", e)
					}
				}
			}()
			l := &LocalExecutor{}
			l.ExecLocal(tt.args.currentDir, tt.args.commands, tt.args.fn)
		})
	}
}
