package getngrokurl

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func GetngrokURL(curlPath string) (string, error) {
	out, err := exec.Command(curlPath, "-s", "http://127.0.0.1:4040/api/tunnels").Output()
	if err != nil {
		return "", err
	}
	output := out[:]
	return processRegexp(string(output)), nil
}

func processRegexp(output string) string {
	str := ""

	re := regexp.MustCompile(`"public_url":"https://([^"]+)"`)
	reurl := regexp.MustCompile(`"https://([^"]+)"`)

	if len(re.FindStringIndex(output)) > 0 {
		str = re.FindString(output)

		if len(reurl.FindStringIndex(str)) > 0 {
			addr := reurl.FindString(str)
			addr = strings.ReplaceAll(addr, "\"", "")
			fmt.Println(addr)
			os.Exit(0)
		}
	}

	return ""
}
