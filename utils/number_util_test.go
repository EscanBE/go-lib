package utils

import (
	"fmt"
	"github.com/EscanBE/go-lib/test_utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
)

func TestConditionalInt64(t *testing.T) {
	tests := []struct {
		name       string
		expression bool
		whenTrue   int64
		whenFalse  int64
		want       int64
	}{
		{
			name:       "true",
			expression: true,
			whenTrue:   99,
			whenFalse:  9,
			want:       99,
		},
		{
			name:       "false",
			expression: false,
			whenTrue:   99,
			whenFalse:  9,
			want:       9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConditionalInt64(tt.expression, tt.whenTrue, tt.whenFalse); got != tt.want {
				t.Errorf("ConditionalInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConditionalInt(t *testing.T) {
	tests := []struct {
		name       string
		expression bool
		whenTrue   int
		whenFalse  int
		want       int
	}{
		{
			name:       "true",
			expression: true,
			whenTrue:   99,
			whenFalse:  9,
			want:       99,
		},
		{
			name:       "false",
			expression: false,
			whenTrue:   99,
			whenFalse:  9,
			want:       9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConditionalInt(tt.expression, tt.whenTrue, tt.whenFalse); got != tt.want {
				t.Errorf("ConditionalInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidHexNumber(t *testing.T) {
	t.Run("random", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			num := rand.Int63()
			var hex string
			switch rand.Int() % 3 {
			case 0:
				hex = fmt.Sprintf("%x", num)
				break
			case 1:
				hex = fmt.Sprintf("x%x", num)
				break
			default:
				hex = fmt.Sprintf("0x%x", num)
				break
			}
			expect := num >= 0 && !strings.HasPrefix(hex, "x")
			if !assert.Equal(t, expect, IsValidHexNumber(hex)) {
				break
			}
		}
	})

	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "",
			want:  false,
		},
		{
			input: "invalid",
			want:  false,
		},
		{
			input: "F",
			want:  true,
		},
		{
			input: "G",
			want:  false,
		},
		{
			input: "f",
			want:  true,
		},
		{
			input: "g",
			want:  false,
		},
		{
			input: "0xf",
			want:  true,
		},
		{
			input: "0xg",
			want:  false,
		},
		{
			input: "0x01234567890AbCDEF00",
			want:  true,
		},
		{
			input: "0x01234567890ABCDEFg0",
			want:  false,
		},
		{
			input: "0Xf",
			want:  true,
		},
		{
			input: "xf",
			want:  false,
		},
		{
			input: "AxfF",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsValidHexNumber(tt.input), "IsValidHexNumber(%v)", tt.input)
		})
	}
}

func TestConvertFromHexNumberStringToDecimalString(t *testing.T) {
	t.Run("random", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			num := rand.Int63()
			if num < 0 {
				i--
				continue
			}
			var hex string
			switch rand.Int() % 2 {
			case 0:
				hex = fmt.Sprintf("%x", num)
				break
			default:
				hex = fmt.Sprintf("0x%x", num)
				break
			}
			got, err := convertFromHexNumberStringToDecimalString(hex, true)
			if !assert.Nil(t, err) {
				break
			}
			if !assert.Equalf(t, fmt.Sprintf("%d", num), got, "input [%s], result [%s]", hex, got) {
				break
			}
		}
	})

	tests := []struct {
		input            string
		want             string
		wantErrMsg       string
		bypassValidation bool
	}{
		{
			input:      "0xG",
			want:       "",
			wantErrMsg: "not a valid hex",
		},
		{
			input: "0XF",
			want:  "15",
		},
		{
			input: "0XFF",
			want:  "255",
		},
		{
			input: "FF",
			want:  "255",
		},
		{
			input: "0x0FFff",
			want:  "65535",
		},
		{
			input:      "-F",
			wantErrMsg: "support positive number only",
		},
		{
			input:      "xFF",
			wantErrMsg: "not a valid hex",
			// bypassValidation: true, // if bypass, big.Int can still parse this
		},
		{
			input:      "AxfF",
			wantErrMsg: "not a valid hex",
			// bypassValidation: true, // if bypass, big.Int can still parse this
		},
		{
			input:            "AxfFAg",
			wantErrMsg:       "unable to set hex number",
			bypassValidation: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var got string
			var err error
			if tt.bypassValidation {
				got, err = convertFromHexNumberStringToDecimalString(tt.input, true)
			} else {
				got, err = ConvertFromHexNumberStringToDecimalString(tt.input)
			}
			wantErr := len(tt.wantErrMsg) > 0
			if (err != nil) != wantErr {
				t.Errorf("got err %v (%s), want %v", err, got, wantErr)
				return
			}
			if err != nil && !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsg) {
				return
			}
			assert.Equalf(t, tt.want, got, "got %s, want %s", got, tt.want)
		})
	}
}
