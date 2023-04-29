package hascurl_test

import (
	"localhost/ngtools/internal/hascurl"
	"testing"
)

func TestHasCurl(t *testing.T) {
	path, err := hascurl.HasCurl()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if path == "" {
		t.Errorf("Expected a non-empty string for curl path, but got empty string")
	}
}
