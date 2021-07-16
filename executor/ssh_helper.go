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
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const cmdDefaultTimeout = 60 * time.Second

// SSHConfig for ssh proxy config
// SSHConfig Contains main authority information.
// User field should be a name of user on remote server (ex. john in ssh john@example.com).
// Server field should be a remote machine address (ex. example.com in ssh john@example.com)
// Key is a path to private key on your local machine.
// Port is SSH server port on remote machine.
// Note: easyssh looking for private key in user's home directory (ex. /home/john + Key).
// Then ensure your Key begins from '/' (ex. /.ssh/id_rsa)
type SSHConfig struct {
	User         string
	Server       string
	Key          string
	KeyPath      string
	Port         string
	Passphrase   string
	Password     string
	Timeout      time.Duration
	Ciphers      []string
	KeyExchanges []string
	Fingerprint  string

	// Enable the use of insecure ciphers and key exchange methods.
	// This enables the use of the the following insecure ciphers and key exchange methods:
	// - aes128-cbc
	// - aes192-cbc
	// - aes256-cbc
	// - 3des-cbc
	// - diffie-hellman-group-exchange-sha256
	// - diffie-hellman-group-exchange-sha1
	// Those algorithms are insecure and may allow plaintext data to be recovered by an attacker.
	UseInsecureCipher bool
	Proxy             *SSHConfig
}

// returns ssh.Signer from user you running app home path + cutted key path.
// (ex. pubkey,err := getKeyFile("/.ssh/id_rsa") )
func getKeyFile(keyPath, passphrase string) (ssh.Signer, error) {
	var pubkey ssh.Signer
	var err error
	buf, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	if passphrase != "" {
		pubkey, err = ssh.ParsePrivateKeyWithPassphrase(buf, []byte(passphrase))
	} else {
		pubkey, err = ssh.ParsePrivateKey(buf)
	}

	if err != nil {
		return nil, err
	}

	return pubkey, nil
}

// returns *ssh.ClientConfig and io.Closer.
// if io.Closer is not nil, io.Closer.Close() should be called when
// *ssh.ClientConfig is no longer used.
func getSSHConfig(config *SSHConfig) (*ssh.ClientConfig, io.Closer) {
	var sshAgent io.Closer

	// auths holds the detected ssh auth methods
	var auths []ssh.AuthMethod

	// figure outChan what auths are requested, what is supported
	if config.Password != "" {
		auths = append(auths, ssh.Password(config.Password))
	}
	if config.KeyPath != "" {
		if pubkey, err := getKeyFile(config.KeyPath, config.Passphrase); err != nil {
			log.Printf("getKeyFile error: %v\n", err)
		} else {
			auths = append(auths, ssh.PublicKeys(pubkey))
		}
	}

	if config.Key != "" {
		var signer ssh.Signer
		var err error
		if config.Passphrase != "" {
			signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(config.Key), []byte(config.Passphrase))
		} else {
			signer, err = ssh.ParsePrivateKey([]byte(config.Key))
		}

		if err != nil {
			log.Printf("ssh.ParsePrivateKey: %v\n", err)
		} else {
			auths = append(auths, ssh.PublicKeys(signer))
		}
	}

	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers))
	}

	c := ssh.Config{}
	if config.UseInsecureCipher {
		c.SetDefaults()
		c.Ciphers = append(c.Ciphers, "aes128-cbc", "aes192-cbc", "aes256-cbc", "3des-cbc")
		c.KeyExchanges = append(c.KeyExchanges, "diffie-hellman-group-exchange-sha1", "diffie-hellman-group-exchange-sha256")
	}

	if len(config.Ciphers) > 0 {
		c.Ciphers = append(c.Ciphers, config.Ciphers...)
	}

	if len(config.KeyExchanges) > 0 {
		c.KeyExchanges = append(c.KeyExchanges, config.KeyExchanges...)
	}

	hostKeyCallback := ssh.InsecureIgnoreHostKey()
	if config.Fingerprint != "" {
		hostKeyCallback = func(hostname string, remote net.Addr, publicKey ssh.PublicKey) error {
			if ssh.FingerprintSHA256(publicKey) != config.Fingerprint {
				return fmt.Errorf("ssh: host key fingerprint mismatch")
			}
			return nil
		}
	}

	return &ssh.ClientConfig{
		Config:          c,
		Timeout:         config.Timeout,
		User:            config.User,
		Auth:            auths,
		HostKeyCallback: hostKeyCallback,
	}, sshAgent
}

// Connect to remote server using SSHConfig struct and returns *ssh.Session
func (sshConf *SSHConfig) connect() (*ssh.Client, error) {
	targetConfig, closer := getSSHConfig(sshConf)
	if closer != nil {
		defer func(closer io.Closer) {
			_ = closer.Close()
		}(closer)
	}

	// Enable proxy command
	var client *ssh.Client
	var err error
	if sshConf.Proxy != nil {
		proxyConfig, closer := getSSHConfig(sshConf.Proxy)
		if closer != nil {
			defer func(closer io.Closer) {
				_ = closer.Close()
			}(closer)
		}

		proxyClient, err := ssh.Dial("tcp", net.JoinHostPort(sshConf.Proxy.Server, sshConf.Proxy.Port), proxyConfig)
		if err != nil {
			return nil, err
		}

		conn, err := proxyClient.Dial("tcp", net.JoinHostPort(sshConf.Server, sshConf.Port))
		if err != nil {
			return nil, err
		}

		ncc, channel, req, err := ssh.NewClientConn(conn, net.JoinHostPort(sshConf.Server, sshConf.Port), targetConfig)
		if err != nil {
			return nil, err
		}

		client = ssh.NewClient(ncc, channel, req)
	} else {
		client, err = ssh.Dial("tcp", net.JoinHostPort(sshConf.Server, sshConf.Port), targetConfig)
	}

	return client, err
}

// stream returns one channel that combines the stdout and stderr of the command
// as it is run on the remote machine, and another that sends true when the
// command is done. The sessions and channels will then be closed.
func (sshConf *SSHConfig) stream(command string, timeout ...time.Duration) (<-chan string, <-chan string, <-chan bool, <-chan error, error) {
	// continuously send the command's output over the channel
	stdoutChan := make(chan string)
	stderrChan := make(chan string)
	doneChan := make(chan bool)
	errChan := make(chan error)

	// connect to remote host
	client, err := sshConf.connect()
	if err != nil {
		return stdoutChan, stderrChan, doneChan, errChan, err
	}
	session, err := client.NewSession()
	if err != nil {
		return stdoutChan, stderrChan, doneChan, errChan, err
	}
	// defer session.Close()
	// connect to both outputs (they are of type io.Reader)
	outReader, err := session.StdoutPipe()
	if err != nil {
		_ = client.Close()
		_ = session.Close()
		return stdoutChan, stderrChan, doneChan, errChan, err
	}
	errReader, err := session.StderrPipe()
	if err != nil {
		_ = client.Close()
		_ = session.Close()
		return stdoutChan, stderrChan, doneChan, errChan, err
	}
	err = session.Start(command)
	if err != nil {
		_ = client.Close()
		_ = session.Close()
		return stdoutChan, stderrChan, doneChan, errChan, err
	}

	// combine outputs, create a line-by-line scanner
	stdoutReader := io.MultiReader(outReader)
	stderrReader := io.MultiReader(errReader)
	stdoutScanner := bufio.NewScanner(stdoutReader)
	stderrScanner := bufio.NewScanner(stderrReader)

	go func(stdoutScanner, stderrScanner *bufio.Scanner, stdoutChan, stderrChan chan string, doneChan chan bool, errChan chan error) {
		defer close(doneChan)
		defer close(errChan)
		defer func(client *ssh.Client) {
			_ = client.Close()
		}(client)
		defer func(session *ssh.Session) {
			_ = session.Close()
		}(session)

		// default timeout value
		executeTimeout := cmdDefaultTimeout
		if len(timeout) > 0 {
			executeTimeout = timeout[0]
		}
		timeoutChan := time.After(executeTimeout)
		res := make(chan struct{}, 1)
		var resWg sync.WaitGroup
		resWg.Add(2)

		go func() {
			defer close(stdoutChan)
			for stdoutScanner.Scan() {
				stdoutChan <- stdoutScanner.Text()
			}
			resWg.Done()
		}()

		go func() {
			defer close(stderrChan)
			for stderrScanner.Scan() {
				stderrChan <- stderrScanner.Text()
			}
			resWg.Done()
		}()

		go func() {
			resWg.Wait()
			// close all of our open resources
			res <- struct{}{}
		}()

		select {
		case <-res:
			errChan <- session.Wait()
			doneChan <- true
		case <-timeoutChan:
			stderrChan <- "run Command Timeout!"
			errChan <- nil
			doneChan <- false
		}
	}(stdoutScanner, stderrScanner, stdoutChan, stderrChan, doneChan, errChan)

	return stdoutChan, stderrChan, doneChan, errChan, err
}

// run command on remote machine and returns its stdout as a string
func (sshConf *SSHConfig) run(command string, timeout ...time.Duration) (outStr string, errStr string, isTimeout bool, err error) {
	stdoutChan, stderrChan, doneChan, errChan, err := sshConf.stream(command, timeout...)
	if err != nil {
		return outStr, errStr, isTimeout, err
	}
	// read from the output channel until the done signal is passed
loop:
	for {
		select {
		case isTimeout = <-doneChan:
			break loop
		case outline := <-stdoutChan:
			if outline != "" {
				outStr += outline + "\n"
			}
		case errLine := <-stderrChan:
			if errLine != "" {
				errStr += errLine + "\n"
			}
		case err = <-errChan:
		}
	}
	// return the concatenation of all signals from the output channel
	return outStr, errStr, isTimeout, err
}
