package hascurl

import (
	"os/exec"
)

func HasCurl() (string, error) {
	var curlPath string
	var err error

	if curlPath, err = exec.LookPath("curl"); err != nil {
		return "", err
	}
	return curlPath, nil
}
