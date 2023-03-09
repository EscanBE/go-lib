package abi

import (
	"encoding/hex"
	"fmt"
	gethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

var contractMethodTemplate = regexp.MustCompile("^\\s*[_aA-zZ\\d]+\\s*\\((((,?)\\s?[_aA-zZ\\d]+\\s*)+)?\\)$")

type SmartContractCallInput struct {
	MethodName string
	MethodID   []byte
	Inputs     []SmartContractCallInputArg
}

type SmartContractCallInputArg struct {
	Index int
	Type  string
	Value interface{}
}

func DecodeSmartContractInputFromTemplate(methodTemplate string, input string) (result SmartContractCallInput, err error) {
	methodTemplate = strings.TrimSpace(methodTemplate)
	if !contractMethodTemplate.MatchString(methodTemplate) {
		return SmartContractCallInput{}, fmt.Errorf("bad method template")
	}

	input = strings.TrimPrefix(strings.TrimSpace(input), "0x")
	inputData, err := hex.DecodeString(input)
	if err != nil {
		return SmartContractCallInput{}, errors.Wrap(err, "failed to decode input")
	}
	inputDataLen := len(inputData)
	if inputDataLen < 4 {
		return SmartContractCallInput{}, fmt.Errorf("bad input length")
	}

	methodName, err := extractMethodName(methodTemplate)
	if err != nil {
		return SmartContractCallInput{}, errors.Wrap(err, "failed to extract method name from method template")
	}

	methodParams, err := extractMethodParams(methodTemplate)
	if err != nil {
		return SmartContractCallInput{}, errors.Wrap(err, "failed to extract method params from method template")
	}

	if len(methodParams) == 0 && len(inputData) > 4 {
		return SmartContractCallInput{}, fmt.Errorf("method contains no params but still have input")
	}

	if len(methodParams) > 0 && len(inputData) == 4 {
		return SmartContractCallInput{}, fmt.Errorf("method has params but empty input")
	}

	abiJson := fmt.Sprintf(`[{"type": "function","name": "%s","inputs": [`, methodName)

	for i, param := range methodParams {
		if i > 0 {
			abiJson += `,`
		}
		abiJson += fmt.Sprintf(`{"internalType": "%s","name": "p%d","type": "%s"}`, param, i+1, param)
	}
	abiJson += `]}]`

	abiSpec, err := gethabi.JSON(strings.NewReader(abiJson))
	if err != nil {
		panic(errors.Wrap(err, "failed to parse abi json"))
	}

	sigData := inputData[:4]
	argData := inputData[4:]

	method, err := abiSpec.MethodById(sigData)
	if err != nil {
		return SmartContractCallInput{}, errors.Wrap(err, "method by id from input does not exists")
	}
	result.MethodName = method.Name
	result.MethodID = method.ID

	if len(methodParams) == 0 && len(inputData) == 4 {
		result.Inputs = make([]SmartContractCallInputArg, 0)
		return result, nil
	}

	inputs, err := method.Inputs.Unpack(argData)
	if err != nil {
		return SmartContractCallInput{}, errors.Wrap(err, "failed to unpack input arguments")
	}

	result.Inputs = make([]SmartContractCallInputArg, len(inputs))
	for i, input := range inputs {
		result.Inputs[i] = SmartContractCallInputArg{
			Index: i,
			Type:  methodParams[i],
			Value: input,
		}
	}

	return result, nil
}

func extractMethodName(methodTemplate string) (methodName string, err error) {
	idx := strings.Index(methodTemplate, "(")
	if idx < 0 {
		return "", fmt.Errorf("bad method template")
	}
	if idx == 0 {
		return "", fmt.Errorf("missing method name")
	}
	return strings.TrimSpace(methodTemplate[:idx]), nil
}

func extractMethodParams(methodTemplate string) (methodParams []string, err error) {
	openBIdx := strings.Index(methodTemplate, "(")

	if openBIdx < 0 {
		return nil, fmt.Errorf("bad method template")
	}

	methodTemplate = strings.TrimSpace(strings.TrimSuffix(methodTemplate, ")"))

	result := make([]string, 0)

	if openBIdx < len(methodTemplate)-1 {
		paramsPart := methodTemplate[openBIdx+1:]
		spl := strings.Split(paramsPart, ",")
		for _, paramPart := range spl {
			normalized := strings.TrimSpace(paramPart)
			if len(normalized) < 1 {
				return nil, fmt.Errorf("empty param pattern")
			}
			spl2 := strings.Split(normalized, ",")
			for _, candidate := range spl2 {
				if len(candidate) > 0 {
					result = append(result, candidate)
					break
				}
			}
		}
	}

	return result, nil
}

func Get4byteMethodId(methodName string, methodParams []string) (string, error) {
	if len(methodName) < 1 {
		return "", fmt.Errorf("method name is empty")
	}

	signature := methodName
	for i, param := range methodParams {
		if len(param) < 1 {
			return "", fmt.Errorf("param content can not be empty")
		}
		if i == 0 {
			signature += "("
		} else {
			signature += ","
		}
		signature += param
	}
	signature += ")"

	bytes := crypto.Keccak256([]byte(signature))

	return hex.EncodeToString(bytes[0:4]), nil
}
