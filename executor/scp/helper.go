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
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type ResponseType = uint8

const (
	Ok      ResponseType = 0
	Warning ResponseType = 1
	Error   ResponseType = 2
)

// Response There are tree types of responses that the remote can send back:
// ok, warning and error
//
// The difference between warning and error is that the connection is not closed by the remote,
// however, a warning can indicate a file transfer failure (such as invalid destination directory)
// and such be handled as such.
//
// All responses except for the `Ok` type always have a message (although these can be empty)
//
// The remote sends a confirmation after every SCP command, because a failure can occur after every
// command, the response should be read and checked after sending them.
type Response struct {
	Type    ResponseType
	Message string
}

// Checks the response it reads from the remote, and will return a single error in case
// of failure
func checkResponse(r io.Reader) error {
	response, err := ParseResponse(r)
	if err != nil {
		return err
	}

	if response.IsFailure() {
		return errors.New(response.GetMessage())
	}

	return nil

}

// CopyN is an adaptation of io.CopyN that keeps reading if it did not return
// a sufficient amount of bytes.
func CopyN(writer io.Writer, src io.Reader, size int64) (int64, error) {
	var total int64
	total = 0
	for total < size {
		n, err := io.CopyN(writer, src, size)
		if err != nil {
			return 0, err
		}
		total += n
	}

	return total, nil
}

// ParseResponse Reads from the given reader (assuming it is the output of the remote) and parses it into a Response structure
func ParseResponse(reader io.Reader) (Response, error) {
	buffer := make([]uint8, 1)
	_, err := reader.Read(buffer)
	if err != nil {
		return Response{}, err
	}

	responseType := buffer[0]
	message := ""
	if responseType > 0 {
		bufferedReader := bufio.NewReader(reader)
		message, err = bufferedReader.ReadString('\n')
		if err != nil {
			return Response{}, err
		}
	}

	return Response{responseType, message}, nil
}

func (r *Response) IsOk() bool {
	return r.Type == Ok
}

func (r *Response) IsWarning() bool {
	return r.Type == Warning
}

// IsError Returns true when the remote responded with an error
func (r *Response) IsError() bool {
	return r.Type == Error
}

// IsFailure Returns true when the remote answered with a warning or an error
func (r *Response) IsFailure() bool {
	return r.Type > 0
}

// GetMessage Returns the message the remote sent back
func (r *Response) GetMessage() string {
	return r.Message
}

type FileInfos struct {
	Message     string
	Filename    string
	Permissions string
	Size        int64
}

func (r *Response) ParseFileInfos() (*FileInfos, error) {
	message := strings.ReplaceAll(r.Message, "\n", "")
	parts := strings.Split(message, " ")
	if len(parts) < 3 {
		return nil, errors.New("unable to parse message as file infos")
	}

	size, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &FileInfos{
		Message:     r.Message,
		Permissions: parts[0],
		Size:        int64(size),
		Filename:    parts[2],
	}, nil
}

// Ack Writes an `Ack` message to the remote, does not await its response, a separate call to ParseResponse is
// therefore required to check if the acknowledgement succeeded
func Ack(writer io.Writer) error {
	var msg = []byte{0}
	n, err := writer.Write(msg)
	if err != nil {
		return err
	}
	if n < len(msg) {
		return errors.New("Failed to write ack buffer")
	}
	return nil
}
