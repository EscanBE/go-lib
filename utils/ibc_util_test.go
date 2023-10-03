package utils

import (
	"fmt"
	"testing"
)

func TestBuildIbcDenom(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		path      string
		baseDenom string
		want      string
	}{
		{
			path:      "transfer/channel-0",
			baseDenom: "uosmo",
			want:      "ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518",
		},
		{
			path:      "transfer/channel-3",
			baseDenom: "uatom",
			want:      "ibc/A4DB47A9D3CF9A068D454513891B526702455D3EF08FB9EB558C561F9DC2B701",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s/%s", tt.path, tt.baseDenom), func(t *testing.T) {
			if got := BuildIbcDenom(tt.path, tt.baseDenom); got != tt.want {
				t.Errorf("BuildIbcDenom() = %v, want %v", got, tt.want)
			}
		})
	}
}
