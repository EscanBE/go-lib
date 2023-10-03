package utils

import (
	"fmt"
	"strings"
)

// BuildIbcDenom returns IBC denom based on algorithm provided by IBC docs: Sha256(path/baseDenom)
func BuildIbcDenom(path, baseDenom string) string {
	return fmt.Sprintf("ibc/%s", strings.ToUpper(Sha256(fmt.Sprintf("%s/%s", path, baseDenom))))
}
