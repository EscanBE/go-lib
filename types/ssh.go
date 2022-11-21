package types

import (
	"fmt"
	"github.com/EscanBE/go-lib/validation"
	"net"
	"reflect"
	"strconv"
)

// === SshRemoteServer

// SshRemote holds target endpoint information
type SshRemote struct {
	Host             string
	Port             string
	Username         string
	AuthByPrivateKey *SshAuthPrivateKey
}

// NewSshRemoteEndpoint returns a SshRemote instance with endpoint info and supplied auth type if any
func NewSshRemoteEndpoint(host, port, username string, authType interface{}) (*SshRemote, error) {
	if len(host) < 1 {
		return nil, fmt.Errorf("host is blank")
	}

	if len(port) < 1 {
		port = "22"
	} else {
		portNo, err := strconv.ParseInt(port, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid port number [%s]: unable to parse", port)
		}
		if !validation.IsValidPort(int(portNo)) {
			return nil, fmt.Errorf("invalid port number [%s]: out of range", port)
		}
	}

	if len(username) < 1 {
		return nil, fmt.Errorf("username is blank")
	}

	if authType == nil {
		return nil, fmt.Errorf("at least one authentication type is required")
	}

	result := SshRemote{
		Host:     host,
		Port:     port,
		Username: username,
	}

	if pk, ok1 := authType.(SshAuthPrivateKey); ok1 {
		result.AuthByPrivateKey = &pk
	} else {
		return nil, fmt.Errorf("not supported auth type %s", reflect.TypeOf(authType))
	}

	return &result, nil
}

// WithAuthByPrivateKey supplies auth type of private key for the remote endpoint
func (sr *SshRemote) WithAuthByPrivateKey(authPk SshAuthPrivateKey) {
	sr.AuthByPrivateKey = &authPk
}

// GetEndpoint returns endpoint as combined of host and port
func (sr *SshRemote) GetEndpoint() string {
	if len(sr.Port) < 1 {
		return sr.Host
	}
	return net.JoinHostPort(sr.Host, sr.Port)
}

// IsAuthByPrivateKey checks if private key authentication was supplied (regardless success/valid or not)
func (sr *SshRemote) IsAuthByPrivateKey() bool {
	return sr.AuthByPrivateKey != nil
}

// === SshAuthPrivateKey

// SshAuthPrivateKey holds target endpoint auth information
type SshAuthPrivateKey struct {
	PrivateKey []byte
	Passphrase []byte
}

// NewSshAuthByPrivateKey returns an instance which holds target endpoint auth information
func NewSshAuthByPrivateKey(privateKey, passphrase []byte) SshAuthPrivateKey {
	return SshAuthPrivateKey{
		PrivateKey: privateKey,
		Passphrase: passphrase,
	}
}
