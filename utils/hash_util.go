package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

// Keccak256Hash computes and returns Keccak-256 value of an input string
func Keccak256Hash(input string) string {
	return hex.EncodeToString(crypto.Keccak256Hash([]byte(input)).Bytes())
}

// Sha256 returns SHA256 checksum string value of an input string
func Sha256(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}
