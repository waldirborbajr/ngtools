package startngrok

import (
	"fmt"
	"localhost/ngtools/internal/secureexec"
	"os"
	"time"
)

func StartNGRok(protocol string, port string) {

	ngrokArgs := []string{"-c", protocol, port}
	cmd := secureexec.Command("/usr/bin/bash", ngrokArgs...)
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(stdout))

	time.Sleep(2 * time.Second)
}
