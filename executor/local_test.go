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
	"io/fs"
	"os"
	"path/filepath"
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

func Test_copyOrMove(t *testing.T) {
	type args struct {
		src  string
		dest string
		perm fs.FileMode
		move bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "copy file",
			args: args{
				src:  "test_copy_file",
				dest: "dest_dir",
				perm: 0644,
				move: false,
			},
			wantErr: false,
		},
		{
			name: "copy file with other permission",
			args: args{
				src:  "test_copy_file",
				dest: "dest_dir",
				perm: 0744,
				move: false,
			},
			wantErr: false,
		},
		{
			name: "move file",
			args: args{
				src:  "test_copy_file",
				dest: "dest_dir",
				perm: 0644,
				move: true,
			},
			wantErr: false,
		},
		{
			name: "move file with other permission",
			args: args{
				src:  "test_copy_file",
				dest: "dest_dir",
				perm: 0744,
				move: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpSrc, err := os.MkdirTemp("", "")
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}

			tmpDest, err := os.MkdirTemp("", "")
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}

			src := filepath.Join(tmpSrc, tt.args.src)
			dest := filepath.Join(tmpDest, tt.args.dest)
			file, _ := os.OpenFile(src, os.O_RDONLY|os.O_CREATE, 0644)
			if err = file.Close(); (err != nil) != tt.wantErr {
				t.Errorf("copyOrMove() error = %v, wantErr %v", err, tt.wantErr)
			}

			err = os.MkdirAll(dest, 0777)
			if (err != nil) != tt.wantErr {
				t.Errorf("copyOrMove() error = %v, wantErr %v", err, tt.wantErr)
			}

			// run test
			if err = copyOrMove(src, filepath.Join(dest, tt.args.src), tt.args.perm, tt.args.move); (err != nil) != tt.wantErr {
				t.Errorf("copyOrMove() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				info, err := os.Stat(filepath.Join(dest, tt.args.src))
				if err != nil || os.IsNotExist(err) {
					t.Errorf("copyOrMove() error = %v, wantErr %v", err, tt.wantErr)
				}
				if info.Mode() != tt.args.perm {
					t.Errorf("copyOrMove() error = %v, want %v, got %v", "wrong permission", tt.args.perm, info.Mode())
				}

				// check if move
				if tt.args.move {
					info, err = os.Stat(src)
					if (err != nil) != tt.wantErr && !os.IsNotExist(err) {
						t.Errorf("copyOrMove() error = %v; moved file shoud not exists", err)
					}
				}
			}
		})
	}
}

func Test_syncOne(t *testing.T) {
	type args struct {
		src  string
		dest string
		opts *CopyOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "single file with abs src and rel dest",
			args: args{
				src:  "test_copy_file",
				dest: "dest_dir",
				opts: DefaultCopyOpts(),
			},
			wantErr: false,
		},
		{
			name: "single file with abs src and abs dest",
			args: args{
				src:  "test_copy_file",
				dest: "dest_dir",
				opts: DefaultCopyOpts(),
			},
			wantErr: false,
		},
		{
			name: "single file with rel src and abs dest",
			args: args{
				src:  "test_copy_file",
				dest: "dest_dir",
				opts: DefaultCopyOpts(),
			},
			wantErr: false,
		},
		{
			name: "single file with rel src and rel dest",
			args: args{
				src:  "test_copy_file",
				dest: "dest_dir",
				opts: DefaultCopyOpts(),
			},
			wantErr: false,
		},
		{
			name: "folder with abs src and abs dest",
			args: args{
				src:  "src_dir",
				dest: "dest_dir",
				opts: DefaultCopyOpts(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "single file with abs src and rel dest" {
				// prepare data
				tmpSrc, err := os.MkdirTemp("", "")
				if (err != nil) != tt.wantErr {
					t.Error(err)
				}

				src := filepath.Join(tmpSrc, tt.args.src)
				file, _ := os.OpenFile(src, os.O_RDONLY|os.O_CREATE, 0644)
				if err = file.Close(); (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				}

				if err = syncOne(src, tt.args.dest, tt.args.opts); (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				} else {
					cwd, _ := os.Getwd()
					_, err = os.Stat(filepath.Join(cwd, tt.args.dest, tt.args.src))
					if (err != nil) != tt.wantErr {
						t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}
			if tt.name == "single file with abs src and abs dest" {
				// prepare data
				tmpSrc, err := os.MkdirTemp("", "")
				if (err != nil) != tt.wantErr {
					t.Error(err)
				}

				tmpDest, err := os.MkdirTemp("", "")
				if (err != nil) != tt.wantErr {
					t.Error(err)
				}

				src := filepath.Join(tmpSrc, tt.args.src)
				dest := filepath.Join(tmpDest, tt.args.dest)
				file, _ := os.OpenFile(src, os.O_RDONLY|os.O_CREATE, 0644)
				if err = file.Close(); (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				}

				if err = syncOne(src, dest, tt.args.opts); (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				} else {
					_, err = os.Stat(filepath.Join(dest, tt.args.src))
					if (err != nil) != tt.wantErr {
						t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}
			if tt.name == "single file with rel src and abs dest" {
				// prepare data
				tmpSrc, err := os.Getwd()
				if (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				}

				tmpDest, err := os.MkdirTemp("", "")
				if (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				}

				src := filepath.Join(tmpSrc, tt.args.src)
				dest := filepath.Join(tmpDest, tt.args.dest)
				file, _ := os.OpenFile(src, os.O_RDONLY|os.O_CREATE, 0644)
				if err = file.Close(); (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				}
				defer func(name string) {
					_ = os.Remove(name)
				}(src)

				if err = syncOne(src, dest, tt.args.opts); (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				} else {
					_, err = os.Stat(filepath.Join(dest, tt.args.src))
					if (err != nil) != tt.wantErr {
						t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}
			if tt.name == "single file with rel src and rel dest" {
				// prepare data
				tmpSrc, err := os.Getwd()
				if (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				}

				src := filepath.Join(tmpSrc, tt.args.src)
				file, _ := os.OpenFile(src, os.O_RDONLY|os.O_CREATE, 0644)
				if err = file.Close(); (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				}
				defer func(name string) {
					_ = os.Remove(name)
				}(src)

				if err = syncOne(src, tt.args.dest, tt.args.opts); (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				} else {
					dest := filepath.Join(tmpSrc, tt.args.dest, tt.args.src)
					_, err = os.Stat(dest)
					if (err != nil) != tt.wantErr {
						t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
					}
					defer func(name string) {
						_ = os.Remove(name)
						_ = os.Remove(filepath.Dir(name))
					}(dest)
				}
			}
			if tt.name == "folder with abs src and abs dest" {
				// prepare data
				tmpSrc, err := os.MkdirTemp("", "")
				if (err != nil) != tt.wantErr {
					t.Error(err)
				}

				tmpDest, err := os.MkdirTemp("", "")
				if (err != nil) != tt.wantErr {
					t.Error(err)
				}

				src := filepath.Join(tmpSrc, tt.args.src)
				dest := filepath.Join(tmpDest, tt.args.dest)

				// make some folder and files
				err = os.MkdirAll(filepath.Join(src, "test/copy/or/move"), 0755)
				err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.IsDir() {
						filePath := filepath.Join(path, "test_file")
						file, err := os.Create(filePath)
						if err != nil {
							return err
						}
						if err = file.Close(); (err != nil) != tt.wantErr {
							t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
						}
					}
					return nil
				})

				if err = syncOne(src, dest, tt.args.opts); (err != nil) != tt.wantErr {
					t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
				} else {
					for _, p := range []string{"test", "test/copy", "test/copy/or", "test/copy/or/move"} {
						_, err = os.Stat(filepath.Join(dest, p, "test_file"))
						if (err != nil) != tt.wantErr {
							t.Errorf("syncOne() error = %v, wantErr %v", err, tt.wantErr)
						}

					}
				}
			}

		})
	}
}
