package verifyos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifiOSNotWindows(t *testing.T) {
	assert.Equal(t, VerifyOS(), false)

}

func TestVerifiOSIstWindows(t *testing.T) {
	assert.Equal(t, !VerifyOS(), true)

}
