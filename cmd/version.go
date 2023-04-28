/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
)

var version = "0.0.0-dev"

func (cli *Cli) version() {

	// cmd := secureexec.Command("ngrok", "--version")
	// // stdout, _ := cmd.Output()
	// stdout, _ := cmd.CombinedOutput()
	// ngrokVersion := strings.Trim(strings.TrimPrefix(string(stdout), "git version "), " \r\n")

	fmt.Println(version)

}
