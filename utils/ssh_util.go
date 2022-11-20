package utils

import (
	"bytes"
	"fmt"
	"github.com/EscanBE/go-lib/types"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"net"
	"strconv"
)

func ExecuteRemoteCommandViaSSH(remoteCommand, remoteHost, sshPort string, sshAuth types.SshAuthPrivateKey) (string, error) {
	if len(remoteCommand) < 1 || IsBlank(remoteCommand) {
		return "", fmt.Errorf("no command was provided")
	}
	if len(remoteHost) < 1 || IsBlank(remoteHost) {
		return "", fmt.Errorf("no remote host was provided")
	}
	if len(sshPort) < 1 || IsBlank(sshPort) {
		return "", fmt.Errorf("no SSH port was provided")
	}
	portNo, errPort := strconv.ParseInt(sshPort, 10, 32)
	if errPort != nil {
		return "", errors.Wrap(errPort, "supplied SSH port is invalid format")
	}
	if portNo < 1 || portNo > 65535 {
		return "", fmt.Errorf("bad SSH port no.%d", portNo)
	}
	if len(sshAuth.UserName) < 1 || IsBlank(sshAuth.UserName) {
		return "", fmt.Errorf("no username was provided")
	}
	if len(sshAuth.PrivateKey) < 24 {
		return "", fmt.Errorf("bad private key")
	}

	var key ssh.Signer
	var err error
	if len(sshAuth.EncryptionPassword) > 0 {
		key, err = ssh.ParsePrivateKeyWithPassphrase(sshAuth.PrivateKey, sshAuth.EncryptionPassword)
	} else {
		key, err = ssh.ParsePrivateKey(sshAuth.PrivateKey)
	}
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt SSH private key")
	}
	auth := []ssh.AuthMethod{
		ssh.PublicKeys(key),
	}

	// Authentication
	config := &ssh.ClientConfig{
		User: sshAuth.UserName,
		// https://github.com/golang/go/issues/19767
		// as clientConfig is non-permissive by default
		// you can set ssh.InsercureIgnoreHostKey to allow any host
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            auth,
	}
	// Connect
	client, err := ssh.Dial("tcp", net.JoinHostPort(remoteHost, sshPort), config)
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
