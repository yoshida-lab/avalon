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
	"errors"
	"testing"
)

func TestError_String(t *testing.T) {
	type fields struct {
		ErrCode ErrorCode
		Err     error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"IOError",
			fields{
				IOError,
				errors.New("IOError"),
			},
			"IOError (0): IOError",
		},
		{
			"CmdExecutionError",
			fields{
				CmdExecutionError,
				errors.New("CmdExecutionError"),
			},
			"CmdExecutionError (1): CmdExecutionError",
		},
		{
			"LoggingError",
			fields{
				LoggingError,
				errors.New("LoggingError"),
			},
			"LoggingError (2): LoggingError",
		},
		{
			"JoinError",
			fields{
				JoinError,
				errors.New("JoinError"),
			},
			"JoinError (3): JoinError",
		},
		{
			"LocalFileError",
			fields{
				LocalFileError,
				errors.New("LocalFileError"),
			},
			"LocalFileError (4): LocalFileError",
		},
		{
			"RemoteAddrError",
			fields{
				RemoteAddrError,
				errors.New("RemoteAddrError"),
			},
			"RemoteAddrError (5): RemoteAddrError",
		},
		{
			"SSHConnectError",
			fields{
				SSHConnectError,
				errors.New("SSHConnectError"),
			},
			"SSHConnectError (6): SSHConnectError",
		},
		{
			"AuthenticationError",
			fields{
				AuthenticationError,
				errors.New("AuthenticationError"),
			},
			"AuthenticationError (7): AuthenticationError",
		},
		{
			"OtherError",
			fields{
				OtherError,
				errors.New("OtherError"),
			},
			"OtherError (8): OtherError",
		},
		{
			"UNKNOWN",
			fields{
				100,
				errors.New("UNKNOWN"),
			},
			"UNKNOWN (100): UNKNOWN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				errCode: tt.fields.ErrCode,
				err:     tt.fields.Err,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
