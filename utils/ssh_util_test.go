package utils

import "testing"

func TestIsPassphraseCanDecryptPrivateKey(t *testing.T) {
	pk := "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAACmFlczI1Ni1jdHIAAAAGYmNyeXB0AAAAGAAAABBx/sIaS3\nJoJwbYJF6I53+3AAAAEAAAAAEAAAAzAAAAC3NzaC1lZDI1NTE5AAAAILzUyN5j8T+DTm2U\nplS/v/s8DHA7zcYUDIbOJgnCZRl2AAAAkIMPd6Ry7fFv+syyOfqq3wVmjm2JigPCanbAoi\nU3yCbGPLqrs9QLMrv0Ax08b9PijvtsE/+fM9LoCUAGRd+6/+GtPlgo0Xjw4HROO1c1RbjG\nIYipxuuj6lquTD8kOObFtT1zHG1a4jlq7gyDYYDsgfg79QnP8GX0WYF0bqtynU2YLtXYT2\n9DdR/tdsOQoysZAQ==\n-----END OPENSSH PRIVATE KEY-----"
	tests := []struct {
		privateKey []byte
		passphrase string
		want       bool
		wantErr    bool
	}{
		{
			privateKey: []byte(pk),
			passphrase: "123456789",
			want:       true,
			wantErr:    false,
		},
		{
			privateKey: []byte(pk),
			passphrase: "1234567890",
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
