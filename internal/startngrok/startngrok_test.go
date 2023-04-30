package startngrok_test

import (
	"fmt"
	"localhost/ngtools/internal/startngrok"
	"os/exec"
	"strings"
	"testing"
)

func TestStartNGRok(t *testing.T) {
	expectedProtocol := "http"
	expectedPort := "8080"

	startngrok.StartNGRok(expectedProtocol, expectedPort)

	// Verify that ngrok is running with the expected arguments
	// This could be done in various ways, such as:
	// 1. Checking if the ngrok process is running using "ps" command
	// 2. Verifying that ngrok is serving traffic on the expected URL
	// 3. Checking ngrok logs for the expected protocol and port
	// Choose the method that works best for your use case and implement it here.

	// For example, you could check if the ngrok process is running with the expected arguments like this:
	output, err := exec.Command("sh", "-c", "ps -ef | grep ngrok").CombinedOutput()
	if err != nil {
		t.Errorf("Error checking ngrok process: %v", err)
	}

	if !strings.Contains(string(output), fmt.Sprintf("ngrok %s %s", expectedProtocol, expectedPort)) {
		t.Errorf("ngrok process not found with expected protocol and port")
	}
}
