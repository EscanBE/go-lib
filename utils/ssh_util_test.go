package utils

import (
	"github.com/EscanBE/go-lib/test_utils"
	"github.com/EscanBE/go-lib/types"
	"golang.org/x/crypto/ssh"
	"testing"
)

var testPk = "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAACmFlczI1Ni1jdHIAAAAGYmNyeXB0AAAAGAAAABBx/sIaS3\nJoJwbYJF6I53+3AAAAEAAAAAEAAAAzAAAAC3NzaC1lZDI1NTE5AAAAILzUyN5j8T+DTm2U\nplS/v/s8DHA7zcYUDIbOJgnCZRl2AAAAkIMPd6Ry7fFv+syyOfqq3wVmjm2JigPCanbAoi\nU3yCbGPLqrs9QLMrv0Ax08b9PijvtsE/+fM9LoCUAGRd+6/+GtPlgo0Xjw4HROO1c1RbjG\nIYipxuuj6lquTD8kOObFtT1zHG1a4jlq7gyDYYDsgfg79QnP8GX0WYF0bqtynU2YLtXYT2\n9DdR/tdsOQoysZAQ==\n-----END OPENSSH PRIVATE KEY-----"
var passphraseTestPk = "123456789"
var generateRemoteServer = func(privateKey, passphrase []byte) *types.SshRemote {
	rs := types.SshRemote{
		Host:     "192.168.0.1",
		Port:     "22",
		Username: "test",
	}
	if len(privateKey) > 0 && len(passphrase) > 0 {
		rs.AuthByPrivateKey = &types.SshAuthPrivateKey{
			PrivateKey: privateKey,
			Passphrase: passphrase,
		}
	} else if len(privateKey) > 0 {
		rs.AuthByPrivateKey = &types.SshAuthPrivateKey{
			PrivateKey: privateKey,
		}
	}
	return &rs
}

func TestIsPassphraseCanDecryptPrivateKey(t *testing.T) {
	tests := []struct {
		privateKey []byte
		passphrase string
		want       bool
		wantErr    bool
	}{
		{
			privateKey: []byte(testPk),
			passphrase: passphraseTestPk,
			want:       true,
			wantErr:    false,
		},
		{
			privateKey: []byte(testPk),
			passphrase: passphraseTestPk + "-invalid",
			want:       false,
			wantErr:    true,
		},
		{
			privateKey: []byte{},
			passphrase: "",
			want:       false,
			wantErr:    true,
		},
		{
			privateKey: []byte{1},
			passphrase: "",
			want:       false,
			wantErr:    true,
		},
		{
			privateKey: []byte{},
			passphrase: "1",
			want:       false,
			wantErr:    true,
		},
		{
			privateKey: []byte{1},
			passphrase: "1",
			want:       false,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := IsPassphraseCanDecryptPrivateKey(tt.privateKey, []byte(tt.passphrase))
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPassphraseCanDecryptPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPassphraseCanDecryptPrivateKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecuteRemoteCommandViaSSH(t *testing.T) {
	tests := []struct {
		name          string
		remoteCommand string
		remoteServer  *types.SshRemote
		wantErrMsg    string
		wantPanic     bool
	}{
		{
			name:          "reject authentication type",
			remoteCommand: "sudo reboot",
			remoteServer:  generateRemoteServer(nil, nil),
			wantErrMsg:    "not supported authentication type",
		},
		{
			name:          "require remote server",
			remoteCommand: "sudo reboot",
			remoteServer:  nil,
			wantPanic:     true,
		},
		{
			name:          "require correct passphrase for private key",
			remoteCommand: "sudo reboot",
			remoteServer:  generateRemoteServer([]byte(testPk), nil),
			wantErrMsg:    "failed to decrypt SSH private key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer test_utils.DeferWantPanicDepends(t, tt.wantPanic)

			got, err := ExecuteRemoteCommandViaSSH(tt.remoteCommand, tt.remoteServer)
			wantErr := len(tt.wantErrMsg) > 0
			if (err != nil) != wantErr {
				t.Errorf("ExecuteRemoteCommandViaSSH() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsg) {
				return
			}
			if got != "" {
				t.Errorf("ExecuteRemoteCommandViaSSH() blind test expect empty response")
			}
		})
	}
}

func Test_executeRemoteCommandViaSSHUsingPrivateKey(t *testing.T) {
	tests := []struct {
		name          string
		remoteCommand string
		remoteServer  *types.SshRemote
		wantErrMsg    string
	}{
		{
			name:          "require private key",
			remoteCommand: "sudo reboot",
			remoteServer:  generateRemoteServer(nil, nil),
			wantErrMsg:    "not a SSH auth by private key",
		},
		{
			name:          "bad private key",
			remoteCommand: "sudo reboot",
			remoteServer:  generateRemoteServer([]byte{1}, nil),
			wantErrMsg:    "bad private key",
		},
		{
			name:          "require correct passphrase for private key",
			remoteCommand: "sudo reboot",
			remoteServer:  generateRemoteServer([]byte(testPk), []byte(passphraseTestPk+"-invalid")),
			wantErrMsg:    "failed to decrypt SSH private key",
		},
		{
			name:          "require correct passphrase for private key",
			remoteCommand: "sudo reboot",
			remoteServer:  generateRemoteServer([]byte(testPk), nil),
			wantErrMsg:    "failed to decrypt SSH private key",
		},
		{
			name:          "good config except bad command",
			remoteCommand: "",
			remoteServer:  generateRemoteServer([]byte(testPk), []byte(passphraseTestPk)),
			wantErrMsg:    "no command was provided",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executeRemoteCommandViaSSHUsingPrivateKey(tt.remoteCommand, tt.remoteServer)
			wantErr := len(tt.wantErrMsg) > 0
			if (err != nil) != wantErr {
				t.Errorf("executeRemoteCommandViaSSHUsingPrivateKey() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsg) {
				return
			}
			if got != "" {
				t.Errorf("ExecuteRemoteCommandViaSSH() blind test expect empty response")
			}
		})
	}
}

var privateKeyAuth = func() []ssh.AuthMethod {
	key, err := ssh.ParsePrivateKeyWithPassphrase([]byte(testPk), []byte(passphraseTestPk))
	if err != nil {
		return nil
	}
	return []ssh.AuthMethod{
		ssh.PublicKeys(key),
	}
}

func Test_executeRemoteCommandViaSSH(t *testing.T) {
	tests := []struct {
		name          string
		remoteCommand string
		remoteServer  *types.SshRemote
		auth          []ssh.AuthMethod
		want          string
		wantErrMsg    string
	}{
		{
			name:          "require command",
			remoteCommand: "",
			remoteServer:  generateRemoteServer(nil, nil),
			auth:          privateKeyAuth(),
			wantErrMsg:    "no command was provided",
		},
		{
			name:          "require host",
			remoteCommand: "sudo reboot",
			remoteServer: &types.SshRemote{
				Host:     "",
				Port:     "22",
				Username: "test",
			},
			auth:       privateKeyAuth(),
			wantErrMsg: "no remote host was provided",
		},
		{
			name:          "require port",
			remoteCommand: "sudo reboot",
			remoteServer: &types.SshRemote{
				Host:     "192.168.0.1",
				Port:     " \t\r\n",
				Username: "test",
			},
			auth:       privateKeyAuth(),
			wantErrMsg: "no SSH port was provided",
		},
		{
			name:          "bad SSH port",
			remoteCommand: "sudo reboot",
			remoteServer: &types.SshRemote{
				Host:     "192.168.0.1",
				Port:     "0",
				Username: "test",
			},
			auth:       privateKeyAuth(),
			wantErrMsg: "bad SSH port",
		},
		{
			name:          "bad SSH port",
			remoteCommand: "sudo reboot",
			remoteServer: &types.SshRemote{
				Host:     "192.168.0.1",
				Port:     "77777",
				Username: "test",
			},
			auth:       privateKeyAuth(),
			wantErrMsg: "bad SSH port",
		},
		{
			name:          "bad SSH port",
			remoteCommand: "sudo reboot",
			remoteServer: &types.SshRemote{
				Host:     "192.168.0.1",
				Port:     "bad",
				Username: "test",
			},
			auth:       privateKeyAuth(),
			wantErrMsg: "port is invalid format",
		},
		{
			name:          "require username",
			remoteCommand: "sudo reboot",
			remoteServer: &types.SshRemote{
				Host:     "192.168.0.1",
				Port:     "22",
				Username: "",
			},
			auth:       privateKeyAuth(),
			wantErrMsg: "no username was provided",
		},
		{
			name:          "require auth",
			remoteCommand: "sudo reboot",
			remoteServer: &types.SshRemote{
				Host:     "192.168.0.1",
				Port:     "22",
				Username: "test",
			},
			auth:       nil,
			wantErrMsg: "no auth was provided",
		},
		{
			name:          "failed to connect",
			remoteCommand: "sudo reboot",
			remoteServer: &types.SshRemote{
				Host:     "192.168.0.1a9",
				Port:     "22222",
				Username: "test",
			},
			auth:       privateKeyAuth(),
			wantErrMsg: "failed to connect to remote server",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executeRemoteCommandViaSSH(tt.remoteCommand, tt.remoteServer, tt.auth)
			wantErr := len(tt.wantErrMsg) > 0
			if (err != nil) != wantErr {
				t.Errorf("executeRemoteCommandViaSSH() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsg) {
				return
			}
			if got != tt.want {
				t.Errorf("executeRemoteCommandViaSSH() got = %v, want %v", got, tt.want)
			}
		})
	}
}
