/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
	"localhost/ngtools/internal/secureexec"
	"strings"
)

// type startCmd struct{}

func (cli *Cli) start(protocol string, port int) {

	fmt.Println("started... ", protocol, port)

	cmd := secureexec.Command("ngrok", "--version")
	stdout, _ := cmd.Output()
	gitVersion := strings.Trim(strings.TrimPrefix(string(stdout), "git version "), " \r\n")
	fmt.Println(gitVersion)

}
