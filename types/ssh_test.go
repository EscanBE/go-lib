package types

import (
	"reflect"
	"testing"
)

func TestNewSshAuthByPrivateKey(t *testing.T) {
	tests := []struct {
		privateKey []byte
		passphrase []byte
		want       SshAuthPrivateKey
	}{
		{
			privateKey: []byte{1, 2, 3},
			passphrase: []byte{4, 5, 6},
			want: SshAuthPrivateKey{
				PrivateKey: []byte{1, 2, 3},
				Passphrase: []byte{4, 5, 6},
			},
		},
		{
			privateKey: []byte{1, 2, 3},
			want: SshAuthPrivateKey{
				PrivateKey: []byte{1, 2, 3},
			},
		},
		{
			passphrase: []byte{4, 5, 6},
			want: SshAuthPrivateKey{
				Passphrase: []byte{4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NewSshAuthByPrivateKey(tt.privateKey, tt.passphrase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSshAuthByPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSshRemoteEndpoint(t *testing.T) {
	type args struct {
		host     string
		port     string
		username string
		authType interface{}
	}
	tests := []struct {
		args    args
		want    *SshRemote
		wantErr bool
	}{
		{
			args: args{
				host:     "x",
				port:     "80",
				username: "z",
				authType: SshAuthPrivateKey{
					PrivateKey: []byte{1, 2, 3},
				},
			},
			want: &SshRemote{
				Host:     "x",
				Port:     "80",
				Username: "z",
				AuthByPrivateKey: &SshAuthPrivateKey{
					PrivateKey: []byte{1, 2, 3},
				},
			},
			wantErr: false,
		},
		{
			args: args{
				host:     "x",
				port:     "80",
				username: "z",
				authType: 6,
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				host:     "x",
				port:     "",
				username: "z",
				authType: SshAuthPrivateKey{
					PrivateKey: []byte{1, 2, 3},
				},
			},
			want: &SshRemote{
				Host:     "x",
				Port:     "22",
				Username: "z",
				AuthByPrivateKey: &SshAuthPrivateKey{
					PrivateKey: []byte{1, 2, 3},
				},
			},
			wantErr: false,
		},
		{
			args: args{
				host:     "x",
				port:     "70000",
				username: "z",
				authType: SshAuthPrivateKey{
					PrivateKey: []byte{1, 2, 3},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				host:     "x",
				port:     "y",
				username: "z",
				authType: SshAuthPrivateKey{
					PrivateKey: []byte{1, 2, 3},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				host:     "",
				port:     "80",
				username: "z",
				authType: SshAuthPrivateKey{
					PrivateKey: []byte{1, 2, 3},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				host:     "x",
				port:     "80",
				username: "",
				authType: SshAuthPrivateKey{
					PrivateKey: []byte{1, 2, 3},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				host:     "x",
				port:     "80",
				username: "z",
				authType: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				host:     "x",
				port:     "y",
				username: "z",
				authType: args{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := NewSshRemoteEndpoint(tt.args.host, tt.args.port, tt.args.username, tt.args.authType)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSshRemoteEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSshRemoteEndpoint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSshRemote_GetEndpoint(t *testing.T) {
	type fields struct {
		Host string
		Port string
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields: fields{
				Host: "localhost",
				Port: "",
			},
			want: "localhost",
		},
		{
			fields: fields{
				Host: "localhost",
				Port: "22",
			},
			want: "localhost:22",
		},
		{
			fields: fields{
				Host: "",
				Port: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			sr := &SshRemote{
				Host: tt.fields.Host,
				Port: tt.fields.Port,
			}
			if got := sr.GetEndpoint(); got != tt.want {
				t.Errorf("GetEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSshRemote_IsAuthByPrivateKey(t *testing.T) {
	type fields struct {
		AuthByPrivateKey *SshAuthPrivateKey
	}
	tests := []struct {
		fields fields
		want   bool
	}{
		{
			fields: fields{
				AuthByPrivateKey: &SshAuthPrivateKey{},
			},
			want: true,
		},
		{
			fields: fields{
				AuthByPrivateKey: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			sr := &SshRemote{
				AuthByPrivateKey: tt.fields.AuthByPrivateKey,
			}
			if got := sr.IsAuthByPrivateKey(); got != tt.want {
				t.Errorf("IsAuthByPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSshRemote_WithAuthByPrivateKey(t *testing.T) {
	sr := &SshRemote{}
	if sr.IsAuthByPrivateKey() {
		t.Errorf("initialization wrongly")
	}
	sr.WithAuthByPrivateKey(SshAuthPrivateKey{})
	if !sr.IsAuthByPrivateKey() {
		t.Errorf("WithAuthByPrivateKey not working as expected")
	}
}
