package ens

import (
	"strings"
	"testing"
)

func TestUts46ToAscii(t *testing.T) {
	type args struct {
		input                        string
		skipTestConvertBackToUnicode bool
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				input: "evmos.evmos",
			},
			want:    "evmos.evmos",
			wantErr: false,
		},
		{
			args: args{
				input: "Evmos.Org",
			},
			want:    "evmos.org",
			wantErr: false,
		},
		{
			args: args{
				input: "bücher.com",
			},
			want:    "xn--bcher-kva.com",
			wantErr: false,
		},
		{
			args: args{
				input: "παράδειγμα.δοκιμή",
			},
			want:    "xn--hxajbheg2az3al.xn--jxalpdlp",
			wantErr: false,
		},
		{
			args: args{
				input:                        "mycharity。org",
				skipTestConvertBackToUnicode: true,
			},
			want:    "mycharity.org",
			wantErr: false,
		},
		{
			args: args{
				input: "prose ware.com",
			},
			wantErr: true,
		},
		{
			args: args{
				input: "proseware..com",
			},
			wantErr: true,
		},
		{
			args: args{
				input: "my_company.com",
			},
			wantErr: true,
		},
		{
			args: args{
				input: "<script>.evmos",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("_", func(t *testing.T) {
			got, err := Uts46ToAscii(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uts46ToAscii(%s) error = %v (got=%s), wantErr %v", tt.args.input, err, got, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uts46ToAscii(%s) got = %v, want %v", tt.args.input, got, tt.want)
			}
			if err == nil && !tt.args.skipTestConvertBackToUnicode {
				back, err2 := Uts46ToUnicode(got)
				if err2 != nil {
					t.Errorf("Uts46ToUnicode(%s) got err = %v, dont want err", got, err2)
				} else if back != strings.ToLower(tt.args.input) {
					t.Errorf("mis-match after convert back = %s vs input = %s", back, tt.args.input)
				}
			}
		})
	}
}

func TestValidateEnsDomain(t *testing.T) {
	tests := []struct {
		domain    string
		tld       string
		wantValid bool
	}{
		{
			domain:    "victor.evmos",
			tld:       "evmos",
			wantValid: true,
		},
		{
			domain:    "victor.evmos",
			tld:       "eth",
			wantValid: false,
		},
		{
			domain:    "Victor.evmos",
			wantValid: true,
		},
		{
			domain:    "victor..evmos",
			wantValid: false,
		},
		{
			domain:    "victorevmos",
			wantValid: false,
		},
		{
			domain:    "victor.evmos.",
			wantValid: false,
		},
		{
			domain:    "bücher.evmos",
			wantValid: true,
		},
		{
			domain:    "мойдомен.evmos",
			wantValid: true,
		},
		{
			domain:    "παράδειγμα.evmos",
			wantValid: true,
		},
		{
			domain:    "mycharity\u3002evmos",
			wantValid: true,
		},
		{
			domain:    "prose\u0000ware.evmos",
			wantValid: false,
		},
		{
			domain:    "proseware..evmos",
			wantValid: false,
		},
		{
			domain:    "v.evmos",
			wantValid: false,
		},
		{
			domain:    "v.evmos",
			tld:       "evmos",
			wantValid: false,
		},
		{
			domain:    "ab.c",
			wantValid: false,
		},
		{
			domain:    "v.victor.evmos",
			wantValid: true,
		},
		{
			domain:    "heo.evmos",
			wantValid: true,
		},
		{
			domain:    ".heo.evmos",
			wantValid: false,
		},
		{
			domain:    "\u0430.heo.evmos",
			wantValid: true,
		},
		{
			domain:    "\u0061.heo.evmos",
			wantValid: true,
		},
		{
			domain:    "hеllо.evmos",
			wantValid: true,
		},
		{
			domain:    "314159265dd8dbb310642f98f50c066173c1259b.addr.reverse",
			wantValid: false,
		},
		{
			domain:    "314159265dd8dbb310642f98f50c066173c1259b.addr.reverse",
			tld:       "reverse",
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run("_", func(t *testing.T) {
			err := ValidateEnsDomain(tt.domain, tt.tld)
			valid := err == nil
			if valid != tt.wantValid {
				t.Errorf("ValidateEnsDomain(%s) valid = %t, want %t", tt.domain, valid, tt.wantValid)
			}
		})
	}
}

func TestIsValidEnsNode(t *testing.T) {
	tests := []struct {
		node string
		want bool
	}{
		{
			node: "0x311c95021c98874c4d197451be2e141d8af0d380eebd23fea96e677dc2f6910a",
			want: true,
		},
		{
			node: "0x311c95021c98874c4d197451be2e141d8af0d380eebd23fea96e677dc2f691",
			want: false,
		},
		{
			node: "311c95021c98874c4d197451be2e141d8af0d380eebd23fea96e677dc2f6910a",
			want: false,
		},
		{
			node: "311c95021c98874c4d197451be2e141d8af0d380eebd23fea96e677dc2f6910aaa",
			want: false,
		},
		{
			node: "0x311c95021c98874c4d197451be2e141d8af0d380eebd23fea96e677dc2f6910z",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.node, func(t *testing.T) {
			if got := IsValidEnsNode(tt.node); got != tt.want {
				t.Errorf("IsValidEnsNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameHash(t *testing.T) {
	tests := []struct {
		labelOrDomain string
		want          string
		wantErr       bool
	}{
		{
			labelOrDomain: "hеllо.evmos",
			want:          "0x311c95021c98874c4d197451be2e141d8af0d380eebd23fea96e677dc2f6910a",
			wantErr:       false,
		},
		{
			labelOrDomain: "hello.evmos",
			want:          "0x05f114f7e3f585e585577f5300aa152da8070ed628753507da7698fb1f8e7e0a",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run("_", func(t *testing.T) {
			got, err := NameHash(tt.labelOrDomain)
			if (err != nil) != tt.wantErr {
				t.Errorf("NameHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NameHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}
