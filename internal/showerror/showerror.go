package showerror

import (
	"fmt"

	"github.com/fatih/color"
)

func ShowError(msg string) {
	color.Set(color.FgRed)
	fmt.Println(msg)
	color.Unset()
}
