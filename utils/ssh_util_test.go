package utils

import (
	"github.com/EscanBE/go-lib/types"
	"golang.org/x/crypto/ssh"
	"strings"
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
			defer func() {
				r := recover()
				if r == nil && tt.wantPanic {
					t.Errorf("The code did not panic")
				} else if r != nil && !tt.wantPanic {
					t.Errorf("The code should panic")
				}
			}()

			got, err := ExecuteRemoteCommandViaSSH(tt.remoteCommand, tt.remoteServer)
			wantErr := len(tt.wantErrMsg) > 0
			if (err != nil) != wantErr {
				t.Errorf("ExecuteRemoteCommandViaSSH() error = %v, wantErr %v", err, wantErr)
				return
			}
			if err != nil {
				if !strings.Contains(err.Error(), tt.wantErrMsg) {
					t.Errorf("executeRemoteCommandViaSSHUsingPrivateKey() error = [%s], expect contains [%s]", err.Error(), tt.wantErrMsg)
					return
				}
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
			if err != nil {
				if !strings.Contains(err.Error(), tt.wantErrMsg) {
					t.Errorf("executeRemoteCommandViaSSHUsingPrivateKey() error = [%s], expect contains [%s]", err.Error(), tt.wantErrMsg)
					return
				}
			}
			if got != "" {
				t.Errorf("ExecuteRemoteCommandViaSSH() blind test expect empty response")
			}
		})
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executeRemoteCommandViaSSH(tt.remoteCommand, tt.remoteServer, tt.auth)
			wantErr := len(tt.wantErrMsg) > 0
			if (err != nil) != wantErr {
				t.Errorf("executeRemoteCommandViaSSH() error = %v, wantErr %v", err, wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("executeRemoteCommandViaSSH() got = %v, want %v", got, tt.want)
			}
		})
	}
}
