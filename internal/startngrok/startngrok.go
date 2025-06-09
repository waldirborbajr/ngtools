package startngrok

import (
	"fmt"
	"localhost/ngtools/internal/secureexec"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func StartNGRok(protocol string, port string) {
	ngrokArgs := fmt.Sprintf("nohup ngrok %s %s > /dev/null 2> /dev/null < /dev/null &", protocol, port)
	cmd := secureexec.Command("sh", "-c", ngrokArgs)
	cmd.Env = os.Environ()
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	time.Sleep(2 * time.Second)
}

func StartNgrok() error {
	cmd := exec.Command("ngrok", "http", "80")
	if err := cmd.Start(); err != nil {
		return err
	}
	// Salva o PID para uso posterior
	return os.WriteFile("ngrok.pid", []byte(strconv.Itoa(cmd.Process.Pid)), 0600)
}
