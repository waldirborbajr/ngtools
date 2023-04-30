package showerror_test

import (
	"bytes"
	"localhost/ngtools/internal/showerror"
	"testing"

	"github.com/fatih/color"
)

func TestShowError(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf

	msg := "Error message"
	showerror.ShowError(msg)

	// got := buf.String()
	// want := color.RedString(msg + "\n")

	// if got != want {
	// 	t.Errorf("ShowError(%q) = %q; want %q", msg, got, want)
	// }
}
