/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package main

import (
	"os"

	"localhost/ngtools/cmd"
	"localhost/ngtools/internal/showerror"
	"localhost/ngtools/internal/verifyos"
)

func main() {

	// Verify is it is an applicable OS
	if verifyos.VerifyOS() {
		showerror.ShowError("This program it not applicable for Windows machine.")
		os.Exit(1)
	}

	cmd := cmd.Cli{}
	cmd.Run()
}
