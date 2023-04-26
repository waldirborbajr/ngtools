/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
	"localhost/ngtools/internal/secureexec"
	"strings"
)

func (cli *Cli) version() {

	cmd := secureexec.Command("ngrok", "--version")
	// stdout, _ := cmd.Output()
	stdout, _ := cmd.CombinedOutput()
	gitVersion := strings.Trim(strings.TrimPrefix(string(stdout), "git version "), " \r\n")

	fmt.Println(gitVersion)

}
