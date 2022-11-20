package types

type SshAuthPrivateKey struct {
	UserName           string
	PrivateKey         []byte
	EncryptionPassword []byte
}
