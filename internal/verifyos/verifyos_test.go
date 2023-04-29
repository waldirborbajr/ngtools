package verifyos_test

import (
	"localhost/ngtools/internal/verifyos"
	"runtime"
	"strings"
	"testing"
)

func TestVerifyOS(t *testing.T) {
	osName := runtime.GOOS
	expectedResult := strings.EqualFold(strings.ToLower(osName), strings.ToLower("windows"))

	if verifyos.VerifyOS() != expectedResult {
		t.Errorf("VerifyOS() returned wrong result, expected %v but got %v", expectedResult, !expectedResult)
	}
}
