/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
	"localhost/ngtools/internal/secureexec"
	"strings"
)

var (
	// NGTools version
	Version = "0.0.0-dev"
	// GitVersion is the semantic version number.
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate is the build time of ISO8601 format, the output of the $(date -u +'%Y-%m-%dT%H:%M:%SZ') command.
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit is the SHA1 value of Git, the output of the $(git rev-parse HEAD) command.
	GitCommit = "$Format:%H$"
	// GitTreeState Represents the state of the Git repository at build time, with possible values: clean, dirty.
	GitTreeState = ""
)

func (cli *Cli) version() {

	cmd := secureexec.Command("ngrok", "--version")
	// stdout, _ := cmd.Output()
	stdout, _ := cmd.CombinedOutput()
	ngrokVersion := strings.Trim(strings.TrimPrefix(string(stdout), "git version "), " \r\n")

	fmt.Println("NGTools ", Version)
	fmt.Println("ngrok   ", ngrokVersion)
	fmt.Println(" Git Version    ", GitVersion)
	fmt.Println(" Build Date     ", BuildDate)
	fmt.Println(" Git Commmit    ", GitCommit)
	fmt.Println(" Git Tree State ", GitCommit)

}
