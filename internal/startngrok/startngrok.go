package startngrok

import (
	"fmt"
	"localhost/ngtools/internal/secureexec"
	"os"
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
