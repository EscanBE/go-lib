package abi

import (
	"fmt"
	"github.com/EscanBE/go-lib/test_utils"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
	"testing"
)

func Test_extractMethodName(t *testing.T) {
	tests := []struct {
		methodTemplate     string
		wantMethodName     string
		wantErrMsgContains string
	}{
		{
			methodTemplate:     "abc(def)",
			wantMethodName:     "abc",
			wantErrMsgContains: "",
		},
		{
			methodTemplate:     "abc(",
			wantMethodName:     "abc",
			wantErrMsgContains: "",
		},
		{
			methodTemplate:     "(abc",
			wantMethodName:     "",
			wantErrMsgContains: "missing method name",
		},
		{
			methodTemplate:     "(abc(",
			wantMethodName:     "",
			wantErrMsgContains: "missing method name",
		},
		{
			methodTemplate:     "abc",
			wantMethodName:     "",
			wantErrMsgContains: "bad method template",
		},
		{
			methodTemplate:     "",
			wantMethodName:     "",
			wantErrMsgContains: "bad method template",
		},
	}
	for _, tt := range tests {
		t.Run(tt.methodTemplate, func(t *testing.T) {
			gotMethodName, err := extractMethodName(tt.methodTemplate)
			if !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsgContains) {
				return
			}
			if gotMethodName != tt.wantMethodName {
				t.Errorf("extractMethodName() gotMethodName = %v, want %v", gotMethodName, tt.wantMethodName)
			}
		})
	}
}

func Test_extractMethodParams(t *testing.T) {
	tests := []struct {
		methodTemplate     string
		wantMethodParams   []string
		wantErrMsgContains string
	}{
		{
			methodTemplate:     "abc()",
			wantMethodParams:   []string{},
			wantErrMsgContains: "",
		},
		{
			methodTemplate:     "abc(d)",
			wantMethodParams:   []string{"d"},
			wantErrMsgContains: "",
		},
		{
			methodTemplate:     "abc(d,e,f)",
			wantMethodParams:   []string{"d", "e", "f"},
			wantErrMsgContains: "",
		},
		{
			methodTemplate:     "abc(d,e,f",
			wantMethodParams:   []string{"d", "e", "f"},
			wantErrMsgContains: "",
		},
		{
			methodTemplate:     "abc( ",
			wantMethodParams:   []string{},
			wantErrMsgContains: "",
		},
		{
			methodTemplate:     "abc",
			wantErrMsgContains: "bad method template",
		},
		{
			methodTemplate:     "abc(,",
			wantErrMsgContains: "empty param pattern",
		},
	}
	for _, tt := range tests {
		t.Run(tt.methodTemplate, func(t *testing.T) {
			gotMethodParams, err := extractMethodParams(tt.methodTemplate)
			if !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsgContains) {
				return
			}
			if !reflect.DeepEqual(gotMethodParams, tt.wantMethodParams) {
				t.Errorf("extractMethodParams() gotMethodParams = %v, want %v", gotMethodParams, tt.wantMethodParams)
			}
		})
	}
}

func TestDecodeSmartContractInputFromTemplate1(t *testing.T) {
	result, err := DecodeSmartContractInputFromTemplate("mint(uint256)", "0xa0712d680000000000000000000000000000000000000000000000000000000000000052")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Output: %v\n", result)
	}
}

func Test_Get4byteMethodId(t *testing.T) {
	tests := []struct {
		methodName   string
		methodParams []string
		want         string
		wantErr      bool
	}{
		{
			methodName:   "mint",
			methodParams: []string{"uint256"},
			want:         "a0712d68",
			wantErr:      false,
		},
		{
			methodName:   "preMint",
			methodParams: []string{"uint256", "bytes32[]"},
			want:         "5a546223",
			wantErr:      false,
		},
		{
			methodName:   "",
			methodParams: []string{"uint256"},
			want:         "",
			wantErr:      true,
		},
		{
			methodName:   "mint",
			methodParams: []string{""},
			want:         "",
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.methodName, func(t *testing.T) {
			got, err := Get4byteMethodId(tt.methodName, tt.methodParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("get4byteMethodId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("get4byteMethodId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeSmartContractInputFromTemplate(t *testing.T) {
	tests := []struct {
		methodTemplate     string
		input              string
		wantResult         SmartContractCallInput
		wantErrMsgContains string
		wantPanic          bool
	}{
		{
			methodTemplate: "mint(uint256)",
			input:          "0xa0712d680000000000000000000000000000000000000000000000000000000000000052",
			wantResult: SmartContractCallInput{
				MethodName: "mint",
				MethodID:   []byte{0xa0, 0x71, 0x2d, 0x68},
				Inputs: []SmartContractCallInputArg{
					{
						Index: 0,
						Type:  "uint256",
						Value: new(big.Int).SetInt64(82),
					},
				},
			},
		},
		{
			methodTemplate: "mint ( uint256 ) ",
			input:          " a0712d680000000000000000000000000000000000000000000000000000000000000052",
			wantResult: SmartContractCallInput{
				MethodName: "mint",
				MethodID:   []byte{0xa0, 0x71, 0x2d, 0x68},
				Inputs: []SmartContractCallInputArg{
					{
						Index: 0,
						Type:  "uint256",
						Value: new(big.Int).SetInt64(82),
					},
				},
			},
		},
		{
			methodTemplate: "claimRewards(address)",
			input:          "0xef5cfb8c0000000000000000000000000000000000000000000000000000000000000000",
			wantResult: SmartContractCallInput{
				MethodName: "claimRewards",
				MethodID:   []byte{0xef, 0x5c, 0xfb, 0x8c},
				Inputs: []SmartContractCallInputArg{
					{
						Index: 0,
						Type:  "address",
						Value: common.HexToAddress("0x0000000000000000000000000000000000000000"),
					},
				},
			},
		},
		{
			methodTemplate: "swapExactTokensForETH(uint256,uint256,address[],address,uint256)",
			input:          "0x18cbafe500000000000000000000000000000000000000000000000048dbbb89270c824b00000000000000000000000000000000000000000000000001a5e77444bcf28700000000000000000000000000000000000000000000000000000000000000a00000000000000000000000003c87b5eb32a3ff6e9441c75ae1ce47de26e5726300000000000000000000000000000000000000000000000000000000639e0665000000000000000000000000000000000000000000000000000000000000000200000000000000000000000047685b6ac7bb4de761a57828877a7b8254c8b145000000000000000000000000d4949664cd82660aae99bedc034a0dea8a0bd517",
			wantResult: SmartContractCallInput{
				MethodName: "swapExactTokensForETH",
				MethodID:   []byte{0x18, 0xcb, 0xaf, 0xe5},
				Inputs: []SmartContractCallInputArg{
					{
						Index: 0,
						Type:  "uint256",
						Value: new(big.Int).SetInt64(5249995988370489931),
					},
					{
						Index: 1,
						Type:  "uint256",
						Value: new(big.Int).SetInt64(118755451750642311),
					},
					{
						Index: 2,
						Type:  "address[]",
						Value: []common.Address{common.HexToAddress("0x47685B6AC7bB4de761A57828877A7B8254c8B145"), common.HexToAddress("0xD4949664cD82660AaE99bEdc034a0deA8A0bd517")},
					},
					{
						Index: 3,
						Type:  "address",
						Value: common.HexToAddress("0x3c87b5eb32a3Ff6E9441c75AE1ce47de26E57263"),
					},
					{
						Index: 4,
						Type:  "uint256",
						Value: new(big.Int).SetInt64(1671300709),
					},
				},
			},
		},
		{
			methodTemplate: "mint()",
			input:          "0x1249c58b",
			wantResult: SmartContractCallInput{
				MethodName: "mint",
				MethodID:   []byte{0x12, 0x49, 0xc5, 0x8b},
				Inputs:     []SmartContractCallInputArg{},
			},
		},
		{
			methodTemplate:     "badMethodTemplate",
			input:              "0xef5cfb8c0000000000000000000000000000000000000000000000000000000000000000",
			wantErrMsgContains: "bad method template",
		},
		{
			methodTemplate:     "method()",
			input:              "0xBadInput",
			wantErrMsgContains: "failed to decode input",
		},
		{
			methodTemplate:     "method()",
			input:              "0x112233",
			wantErrMsgContains: "bad input length",
		},
		{
			methodTemplate:     "method(,uint256)",
			input:              "0x11223344",
			wantErrMsgContains: "failed to extract method params",
		},
		{
			methodTemplate:     "method()",
			input:              "0x1122334455",
			wantErrMsgContains: "method contains no params but still have input",
		},
		{
			methodTemplate:     "mint()",
			input:              "0x11223344",
			wantErrMsgContains: "method by id from input does not exists",
		},
		{
			methodTemplate:     "method(uint256)",
			input:              "0x11223344",
			wantErrMsgContains: "method has params but empty input",
		},
		{
			methodTemplate:     "mint(uint256)",
			input:              "0xa0712d680000000000000000000000000052",
			wantErrMsgContains: "failed to unpack",
		},
	}
	for _, tt := range tests {
		t.Run(tt.methodTemplate, func(t *testing.T) {
			defer test_utils.DeferWantPanicDepends(t, tt.wantPanic)
			gotResult, err := DecodeSmartContractInputFromTemplate(tt.methodTemplate, tt.input)
			wantErr := len(tt.wantErrMsgContains) > 0
			if (err != nil) != wantErr {
				t.Errorf("DecodeSmartContractInputFromTemplate() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsgContains) {
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("DecodeSmartContractInputFromTemplate() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
