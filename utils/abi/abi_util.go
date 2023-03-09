package abi

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
)

// FromContractResponseStringToAddress convert response from smart contract response into address.
// CONTRACT: must be valid 32 bytes hex with 0x prefix
func FromContractResponseStringToAddress(input string) (common.Address, error) {
	var addr common.Address
	if len(input) != 66 {
		return addr, fmt.Errorf("must be 32 bytes data")
	}

	for i, c := range input[2:27] {
		if i < 24 {
			if c != '0' {
				return addr, fmt.Errorf("first 12 bytes must be zero")
			}
		}
	}

	addr = common.HexToAddress(input)

	return addr, nil
}

// FromContractResponseBufferToAddress convert response from smart contract response into address.
func FromContractResponseBufferToAddress(input []byte) (common.Address, error) {
	var addr common.Address
	if len(input) != 32 {
		return addr, fmt.Errorf("must be 32 bytes data")
	}

	for _, c := range input[:12] {
		if c != 0 {
			return addr, fmt.Errorf("first 12 bytes must be zero")
		}
	}

	addr = common.BytesToAddress(input)

	return addr, nil
}

// FromContractResponseStringToHash convert response from smart contract response into hash.
// CONTRACT: must be valid 32 bytes hex with 0x prefix
func FromContractResponseStringToHash(input string) (common.Hash, error) {
	var hash common.Hash
	if len(input) != 66 {
		return hash, fmt.Errorf("must be 32 bytes data")
	}

	hash = common.HexToHash(input)

	return hash, nil
}

// FromContractResponseBufferToHash convert response from smart contract response into hash.
func FromContractResponseBufferToHash(input []byte) (common.Hash, error) {
	var hash common.Hash
	if len(input) != 32 {
		return hash, fmt.Errorf("must be 32 bytes data")
	}

	hash = common.BytesToHash(input)

	return hash, nil
}

// FromContractResponseStringToString convert response from smart contract response into string.
// CONTRACT: must be valid hex with 0x prefix
func FromContractResponseStringToString(input string) (string, error) {
	if len(input) < 2+32*2*2 {
		return "", fmt.Errorf("must be at least 64 bytes data")
	}

	buffer, err := hex.DecodeString(input[2:])
	if err != nil {
		return "", fmt.Errorf("bad input")
	}

	return FromContractResponseBufferToString(buffer)
}

// FromContractResponseBufferToString convert response from smart contract response into string.
func FromContractResponseBufferToString(buffer []byte) (string, error) {
	if len(buffer) < 64 {
		return "", fmt.Errorf("must be at least 64 bytes data")
	}

	offset := new(big.Int)

	offset = offset.SetBytes(buffer[:32])

	if offset.Int64() != 32 {
		return "", fmt.Errorf("value of first 32 bytes must be 32")
	}

	offset = offset.SetBytes(buffer[32:64])

	strLength := offset.Int64()

	groupsOf32Bytes := int(math.Ceil(float64(strLength) / 32))

	expectedInputBufferLength := 32 + 32 + 32*groupsOf32Bytes
	if len(buffer) != expectedInputBufferLength {
		return "", fmt.Errorf("bad input buffer length, got %d, want %d", len(buffer), expectedInputBufferLength)
	}

	expectedZeroSinceIdx := int(32 + 32 + strLength)
	for i, b := range buffer {
		if i < expectedZeroSinceIdx {
			continue
		}
		if b != 0 {
			return "", fmt.Errorf("un-expected a non-zero bytes at idx %d = %x", i, b)
		}
	}

	bufferOfValue := buffer[64 : 64+strLength]

	return string(bufferOfValue), nil
}
