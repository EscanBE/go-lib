package utils

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

func Keccak256Hash(input string) string {
	return hex.EncodeToString(crypto.Keccak256Hash([]byte(input)).Bytes())
}
