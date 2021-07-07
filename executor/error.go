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
	errCode ErrorCode
	err     error
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
	switch e.errCode {
	case IOError:
		return fmt.Sprintf("IOError (%d): %v", IOError, e.err)
	case CmdExecutionError:
		return fmt.Sprintf("CmdExecutionError (%d): %v", CmdExecutionError, e.err)
	case LoggingError:
		return fmt.Sprintf("LoggingError (%d): %v", LoggingError, e.err)
	case JoinError:
		return fmt.Sprintf("JoinError (%d): %v", JoinError, e.err)
	case LocalFileError:
		return fmt.Sprintf("LocalFileError (%d): %v", LocalFileError, e.err)
	case RemoteAddrError:
		return fmt.Sprintf("RemoteAddrError (%d): %v", RemoteAddrError, e.err)
	case SSHConnectError:
		return fmt.Sprintf("SSHConnectError (%d): %v", SSHConnectError, e.err)
	case AuthenticationError:
		return fmt.Sprintf("AuthenticationError (%d): %v", AuthenticationError, e.err)
	case OtherError:
		return fmt.Sprintf("OtherError (%d): %v", OtherError, e.err)
	default:
		return fmt.Sprintf("UNKNOWN (%d): %v", e.errCode, e.err)
	}
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) Type() ErrorCode {
	return e.errCode
}

func (e *Error) Message() string {
	return fmt.Sprintf("%+v", e.err)
}
