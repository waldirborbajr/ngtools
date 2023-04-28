/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
	"localhost/ngtools/internal/getngrokurl"
	"localhost/ngtools/internal/hascurl"
	"localhost/ngtools/internal/hasnohup"
	"localhost/ngtools/internal/listprocess"
	"localhost/ngtools/internal/showerror"
	"localhost/ngtools/internal/startngrok"
	"os"
	"strconv"
)

func (cli *Cli) start(protocol string, port int) {
	path, err := hascurl.HasCurl()
	if err != nil {
		showerror.ShowError("Curl not found. Please install it first and run it again.\n")
		os.Exit(1)
	}

	// Check for previusly running instance
	if listprocess.HasProcessRunning() {
		showerror.ShowError("ngrok it is running, please stop it first before start new one instance.")
		os.Exit(1)
	}

	// Remove previusly nohup
	hasnohup.HasNoHup()

	// Start ngrok
	startngrok.StartNGRok(protocol, strconv.Itoa(port))

	// Get url generated
	url, err := getngrokurl.GetngrokURL(path)
	if err != nil {
		showerror.ShowError("Error executing curl. Please verify if ngrok it is up and running.\n")
		os.Exit(1)
	}

	fmt.Println(url)
}
