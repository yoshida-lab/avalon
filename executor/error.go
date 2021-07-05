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

import "fmt"

type Error struct {
	ErrCode ErrorCode
	Err     error
}

type ErrorCode uint8

const (
	IOError ErrorCode = iota
	CmdExecutionError
	LoggingError
	JoinError
	LocalFileError
	RemoteAddrError
	SSHConnectError
	AuthenticationError
	OtherError
)

func (e *Error) String() string {
	switch e.ErrCode {
	case IOError:
		return fmt.Sprintf("IOError (%d): %v", IOError, e.Err)
	case CmdExecutionError:
		return fmt.Sprintf("CmdExecutionError (%d): %v", CmdExecutionError, e.Err)
	case LoggingError:
		return fmt.Sprintf("LoggingError (%d): %v", LoggingError, e.Err)
	case JoinError:
		return fmt.Sprintf("JoinError (%d): %v", JoinError, e.Err)
	case LocalFileError:
		return fmt.Sprintf("LocalFileError (%d): %v", LocalFileError, e.Err)
	case RemoteAddrError:
		return fmt.Sprintf("RemoteAddrError (%d): %v", RemoteAddrError, e.Err)
	case SSHConnectError:
		return fmt.Sprintf("SSHConnectError (%d): %v", SSHConnectError, e.Err)
	case AuthenticationError:
		return fmt.Sprintf("AuthenticationError (%d): %v", AuthenticationError, e.Err)
	case OtherError:
		return fmt.Sprintf("OtherError (%d): %v", OtherError, e.Err)
	default:
		return fmt.Sprintf("UNKNOWN (%d): %v", e.ErrCode, e.Err)
	}
}

func (e *Error) Error() string {
	return e.String()
}
