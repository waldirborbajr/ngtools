/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package main

import (
	"fmt"
	"localhost/ngtools/cmd"
	"localhost/ngtools/internal/verifyos"
	"os"
)

func main() {

	if verifyos.VerifyOS() {
		fmt.Println("This program it not applicable for Windows machine.")
		os.Exit(1)
	}

	cmd := cmd.Cli{}
	cmd.Run()
}
