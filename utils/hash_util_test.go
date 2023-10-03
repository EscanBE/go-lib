package utils

import "testing"

func Test256Hashing(t *testing.T) {
	tests := []struct {
		input         string
		wantKeccak256 string
		wantSha256    string
	}{
		{
			input:         "abc",
			wantKeccak256: "4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45",
			wantSha256:    "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
		},
		{
			input:         "abcd",
			wantKeccak256: "48bed44d1bcd124a28c27f343a817e5f5243190d3c52bf347daf876de1dbbf77",
			wantSha256:    "88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589",
		},
		{
			input:         "",
			wantKeccak256: "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
			wantSha256:    "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got1 := Keccak256Hash(tt.input); got1 != tt.wantKeccak256 {
				t.Errorf("Keccak256Hash() = %v, want %v", got1, tt.wantKeccak256)
			}
			if got2 := Sha256(tt.input); got2 != tt.wantSha256 {
				t.Errorf("Sha256() = %v, want %v", got2, tt.wantSha256)
			}
		})
	}
}
