package hascurl

import (
	"localhost/ngtools/internal/showerror"
	"os"
	"os/exec"
)

func HasCurl() string {
	var curlPath string
	var err error

	if curlPath, err = exec.LookPath("curli"); err != nil {
		showerror.ShowError("Curl not found. Please install it first and run it again.\n")
		os.Exit(1)
	}
	return curlPath
}
