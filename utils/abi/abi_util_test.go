package abi

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestFromContractResponseStringToAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    string
		wantErr bool
	}{
		{
			name:    "normal",
			address: "0x000000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			want:    "0xae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			wantErr: false,
		},
		{
			name:    "normal with capital 0x",
			address: "0X000000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			want:    "0xae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			wantErr: false,
		},
		{
			name:    "normal with capital",
			address: "0x000000000000000000000000AE38558e1cccf2d930e7be45de62bcc95c32b0ef",
			want:    "0xae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			wantErr: false,
		},
		{
			name:    "not enough 32 bytes",
			address: "0x000000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0",
			wantErr: true,
		},
		{
			name:    "first 12 bytes must be zero",
			address: "0x000000000000000000000001ae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			wantErr: true,
		},
		{
			name:    "first 12 bytes must be zero",
			address: "0x100000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			wantErr: true,
		},
		{
			name:    "all zero returns empty addr without err",
			address: "0x0000000000000000000000000000000000000000000000000000000000000000",
			want:    "0x0000000000000000000000000000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromContractResponseStringToAddress(tt.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromContractResponseStringToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if !strings.EqualFold(got.String(), tt.want) {
					t.Errorf("FromContractResponseStringToAddress() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.String() != "0x0000000000000000000000000000000000000000" {
					t.Errorf("FromContractResponseStringToAddress() must returns empty addr on error, but got %s", got)
				}
			}

			// test buffer
			buffer, err2 := hex.DecodeString(tt.address[2:])
			if err2 != nil {
				t.Errorf("failed to decode input addr %s into buffer for testing", tt.address)
				t.FailNow()
			}
			addr2, err2 := FromContractResponseBufferToAddress(buffer)
			if err == nil {
				if err2 != nil {
					t.Errorf("FromContractResponseBufferToAddress() expects no error but got %v", err2)
					t.FailNow()
				}

				if addr2 != got {
					t.Errorf("result from buffer method is different with result from string method: got = %v, want = %v", addr2, got)
				}
			} else {
				if err2 == nil {
					t.Errorf("FromContractResponseBufferToAddress() expects error but got no error")
					t.FailNow()
				}

				if addr2.String() != "0x0000000000000000000000000000000000000000" {
					t.Errorf("FromContractResponseBufferToAddress() must returns empty addr on error, but got %s", addr2)
				}
			}
		})
	}
}

func TestFromContractResponseStringToHash(t *testing.T) {
	tests := []struct {
		name    string
		hash    string
		want    string
		wantErr bool
	}{
		{
			name:    "normal",
			hash:    "0x000000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			want:    "0x000000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			wantErr: false,
		},
		{
			name:    "normal with capital 0x",
			hash:    "0X000000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			want:    "0x000000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			wantErr: false,
		},
		{
			name:    "normal with capital",
			hash:    "0x000000000000000000000000AE38558e1cccf2d930e7be45de62bcc95c32b0ef",
			want:    "0x000000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0ef",
			wantErr: false,
		},
		{
			name:    "not enough 32 bytes",
			hash:    "0x000000000000000000000000ae38558e1cccf2d930e7be45de62bcc95c32b0",
			wantErr: true,
		},
		{
			name: "all zero returns empty addr without err",
			hash: "0x0000000000000000000000000000000000000000000000000000000000000000",
			want: "0x0000000000000000000000000000000000000000000000000000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromContractResponseStringToHash(tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromContractResponseStringToHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if !strings.EqualFold(got.String(), tt.want) {
					t.Errorf("FromContractResponseStringToHash() got = %v, want %v", got, tt.want)
				}
			} else {
				if got.String() != "0x0000000000000000000000000000000000000000000000000000000000000000" {
					t.Errorf("FromContractResponseStringToHash() must returns empty hash on error, but got %s", got)
				}
			}

			// test buffer
			buffer, err2 := hex.DecodeString(tt.hash[2:])
			if err2 != nil {
				t.Errorf("failed to decode input addr %s into buffer for testing", tt.hash)
				t.FailNow()
			}
			hash2, err2 := FromContractResponseBufferToHash(buffer)
			if err == nil {
				if err2 != nil {
					t.Errorf("FromContractResponseBufferToHash() expects no error but got %v", err2)
					t.FailNow()
				}

				if hash2 != got {
					t.Errorf("result from buffer method is different with result from string method: got = %v, want = %v", hash2, got)
				}
			} else {
				if err2 == nil {
					t.Errorf("FromContractResponseBufferToHash() expects error but got no error")
					t.FailNow()
				}

				if hash2.String() != "0x0000000000000000000000000000000000000000000000000000000000000000" {
					t.Errorf("FromContractResponseBufferToHash() must returns empty hash on error, but got %s", hash2)
				}
			}
		})
	}
}

func TestFromContractResponseStringToString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "normal",
			input:   "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000d68d0b56c6cd0be2e65766d6f7300000000000000000000000000000000000000",
			want:    "hеllо.evmos",
			wantErr: false,
		},
		{
			name:    "empty",
			input:   "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000",
			want:    "",
			wantErr: false,
		},
		{
			name:    "non-zero after str data",
			input:   "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000d68d0b56c6cd0be2e65766d6f7311000000000000000000000000000000000000",
			wantErr: true,
		},
		{
			name:    "buffer size not match",
			input:   "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000d68d0b56c6cd0be2e65766d6f73000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			wantErr: true,
		},
		{
			name:    "buffer size not match",
			input:   "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001",
			wantErr: true,
		},
		{
			name:    "first 32 bytes value must be 32",
			input:   "0x0000000000000000000000000000000000000000000000000000000000000021000000000000000000000000000000000000000000000000000000000000000d68d0b56c6cd0be2e65766d6f7300000000000000000000000000000000000000",
			wantErr: true,
		},
		{
			name:    "shorter than 64 bytes",
			input:   "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromContractResponseStringToString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromContractResponseStringToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("FromContractResponseStringToString() got = %v, want %v", got, tt.want)
			}
			if err != nil && got != "" {
				t.Errorf("FromContractResponseStringToString() expected empty result when null, but got %s", got)
			}

			// test buffer
			buffer, err2 := hex.DecodeString(tt.input[2:])
			if err2 != nil {
				t.Errorf("failed to decode input addr %s into buffer for testing", tt.input)
				t.FailNow()
			}
			value2, err2 := FromContractResponseBufferToString(buffer)
			if err == nil {
				if err2 != nil {
					t.Errorf("FromContractResponseBufferToString() expects no error but got %v", err2)
					t.FailNow()
				}

				if value2 != got {
					t.Errorf("result from buffer method is different with result from string method: got = %v, want = %v", value2, got)
				}
			} else {
				if err2 == nil {
					t.Errorf("FromContractResponseBufferToString() expects error but got no error")
					t.FailNow()
				}

				if value2 != "" {
					t.Errorf("FromContractResponseBufferToString() must returns empty string on error, but got %s", value2)
				}
			}
		})
	}
}
