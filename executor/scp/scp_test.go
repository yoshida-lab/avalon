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
	"golang.org/x/crypto/ssh"
	"os"
	"testing"
	"time"
)

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(username string, password string, keyCallBack ssh.HostKeyCallback) (ssh.ClientConfig, error) {

	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: keyCallBack,
	}, nil
}

func establishConnection(t *testing.T) *ssh.Client {
	// Use SSH key authentication from the auth package.
	// During testing we ignore the host key, don't to that when you use this.
	clientConfig, _ := PasswordKey("avalon", "test", ssh.InsecureIgnoreHostKey())
	client, err := ssh.Dial("tcp", "127.0.0.1:2244", &clientConfig)
	if err != nil {
		t.Fatalf("Couldn't establish a connection to the remote server: %s", err)
	}

	return client
}

func TestUpload(t *testing.T) {
	type args struct {
		localFilePath string
		remoteDir     string
		permissions   os.FileMode
		timeout       []time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"upload successfully",
			args{
				localFilePath: "../../tests/data/test.csv",
				remoteDir:     "/home/avalon",
				permissions:   0644,
			},
			false,
		},
		{
			"dest dir not exists",
			args{
				localFilePath: "../../tests/data/test.csv",
				remoteDir:     "/home/avalon/not_exists_dir",
				permissions:   0644,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := establishConnection(t)
			session, err := client.NewSession()
			if err != nil {
				t.Fatalf("Couldn't establish a connection to the remote server: %s", err)
			}
			defer func(client *ssh.Client) {
				_ = client.Close()
			}(client)

			if tt.name == "upload successfully" {
				if err = Upload(tt.args.localFilePath, tt.args.remoteDir, tt.args.permissions, session, tt.args.timeout...); (err != nil) != tt.wantErr {
					t.Errorf("Upload() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if tt.name == "dest dir not exists" {
				errMsg := "upload file failed! 'execution error': Process exited with status 1; 'transfer error': EOF"
				err = Upload(tt.args.localFilePath, tt.args.remoteDir, tt.args.permissions, session, tt.args.timeout...)
				if (err != nil) != tt.wantErr {
					t.Errorf("Upload() error = %v, wantErr %v", err, tt.wantErr)
				} else if (err != nil) && err.Error() != errMsg {
					t.Errorf("Upload() error = %v, wantErrMsg %v", err, errMsg)
				}
			}

		})
	}
}

func TestDownload(t *testing.T) {
	type args struct {
		remoteFilePath string
		localDir       string
		permissions    os.FileMode
		timeout        []time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"download",
			args{
				localDir:       "/Users/liuchang",
				remoteFilePath: "/home/avalon/polyimide_prop.csv",
				permissions:    0644,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := establishConnection(t)
			session, err := client.NewSession()
			if err != nil {
				t.Fatalf("Couldn't establish a connection to the remote server: %s", err)
			}
			defer func(client *ssh.Client) {
				_ = client.Close()
			}(client)

			if err := Download(tt.args.remoteFilePath, tt.args.localDir, tt.args.permissions, session, tt.args.timeout...); (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
