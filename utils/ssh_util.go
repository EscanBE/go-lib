package utils

import (
	"bytes"
	"fmt"
	"github.com/EscanBE/go-lib/types"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"strconv"
)

// ExecuteRemoteCommandViaSSH performs execution of remote command on target server,
// it returns the output string of the execution and an error if any
func ExecuteRemoteCommandViaSSH(remoteCommand string, remoteServer *types.SshRemote) (string, error) {
	if remoteServer == nil {
		panic("remote server is required")
	}
	if remoteServer.IsAuthByPrivateKey() {
		return executeRemoteCommandViaSSHUsingPrivateKey(remoteCommand, remoteServer)
	}

	return "", fmt.Errorf("not supported authentication type or none provided")
}

// executeRemoteCommandViaSSHUsingPrivateKey performs execution of remote command on target server,
// using private key authentication method,
// it returns the output string of the execution and an error if any
func executeRemoteCommandViaSSHUsingPrivateKey(remoteCommand string, remoteServer *types.SshRemote) (string, error) {
	if remoteServer.AuthByPrivateKey == nil {
		return "", fmt.Errorf("not a SSH auth by private key (not supplied)")
	}
	if len(remoteServer.AuthByPrivateKey.PrivateKey) < 24 {
		return "", fmt.Errorf("bad private key")
	}

	var key ssh.Signer
	var err error
	if len(remoteServer.AuthByPrivateKey.Passphrase) > 0 {
		key, err = ssh.ParsePrivateKeyWithPassphrase(remoteServer.AuthByPrivateKey.PrivateKey, remoteServer.AuthByPrivateKey.Passphrase)
	} else {
		key, err = ssh.ParsePrivateKey(remoteServer.AuthByPrivateKey.PrivateKey)
	}
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt SSH private key")
	}
	auth := []ssh.AuthMethod{
		ssh.PublicKeys(key),
	}

	return executeRemoteCommandViaSSH(remoteCommand, remoteServer, auth)
}

// executeRemoteCommandViaSSHUsingPrivateKey performs execution of remote command on target server,
// using private key authentication method,
// it returns the output string of the execution and an error if any
func executeRemoteCommandViaSSH(remoteCommand string, remoteServer *types.SshRemote, auth []ssh.AuthMethod) (string, error) {
	if IsBlank(remoteCommand) {
		return "", fmt.Errorf("no command was provided")
	}
	if IsBlank(remoteServer.Host) {
		return "", fmt.Errorf("no remote host was provided")
	}
	if IsBlank(remoteServer.Port) {
		return "", fmt.Errorf("no SSH port was provided")
	}
	portNo, errPort := strconv.ParseInt(remoteServer.Port, 10, 32)
	if errPort != nil {
		return "", errors.Wrap(errPort, "supplied SSH port is invalid format")
	}
	if portNo < 1 || portNo > 65535 {
		return "", fmt.Errorf("bad SSH port no.%d", portNo)
	}
	if IsBlank(remoteServer.Username) {
		return "", fmt.Errorf("no username was provided")
	}

	// Authentication
	config := &ssh.ClientConfig{
		User: remoteServer.Username,
		// https://github.com/golang/go/issues/19767
		// as clientConfig is non-permissive by default
		// you can set ssh.InsercureIgnoreHostKey to allow any host
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            auth,
	}

	// Connect
	client, err := ssh.Dial("tcp", remoteServer.GetEndpoint(), config)
	if err != nil {
		return "", errors.Wrap(err, "failed to connect to remote server")
	}

	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "failed to start a new SSH session")
	}
	//goland:noinspection GoUnhandledErrorResult
	defer session.Close()
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	// you can also pass what gets input to the stdin, allowing you to pipe
	// content from client to server
	//      session.Stdin = bytes.NewBufferString("My input")

	// Finally, run the command
	err = session.Run(remoteCommand)
	return b.String(), err
}

// IsPassphraseCanDecryptPrivateKey returns true if the provided passphrase can decrypt the encrypted privateKey
func IsPassphraseCanDecryptPrivateKey(privateKey, passphrase []byte) (bool, error) {
	_, err := ssh.ParsePrivateKeyWithPassphrase(privateKey, passphrase)
	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}
