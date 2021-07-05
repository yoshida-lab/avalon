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

type CopyOptions struct {
	// Create new directory when not existing.
	CreateDir bool
	// Overwrite existing files if true (default: false).
	Overwrite bool
	// Skip existing files if true (default: false).
	SkipExist bool
	// Buffer size that specifies the amount of bytes to be moved or copied before the progress handler is called. This only affects functions with progress handlers. (default: 64000)
	BufferSize uint
	// Use move semantic instead of copy
	// if set to true, all file/dirs will be deleted after moving. (default: false)
	MoveContent bool
	// Sets levels reading. Set 0 for read all directory folder (default: 0).
	//
	// Warning: Work only for copy operations!
	Depth uint64
	// file permissions
	FileMode uint32
}

func DefaultCopyOpts() CopyOptions {
	return CopyOptions{
		CreateDir:   true,
		Overwrite:   true,
		SkipExist:   true,
		BufferSize:  64000, // 64kb
		MoveContent: false,
		Depth:       0,
		FileMode:    0o644,
	}
}
