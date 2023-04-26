package verifyos

import (
	"runtime"
	"strings"
)

func VerifyOS() bool {
	return strings.EqualFold(strings.ToLower(runtime.GOOS), strings.ToLower("windows"))
}
