/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
	"localhost/ngtools/internal/getngrokurl"
	"localhost/ngtools/internal/hascurl"
	"localhost/ngtools/internal/hasnohup"
	"localhost/ngtools/internal/showerror"
	"localhost/ngtools/internal/startngrok"
	"os"
	"strconv"
)

// type startCmd struct{}

func (cli *Cli) start(protocol string, port int) {
	var path string
	var err error

	if path, err = hascurl.HasCurl(); err != nil {
		showerror.ShowError("Curl not found. Please install it first and run it again.\n")
		os.Exit(1)
	}

	// Remove previusly nohup
	hasnohup.HasNoHup()

	// Start ngrok
	startngrok.StartNGRok(protocol, strconv.Itoa(port))

	// Get url generated
	var url string
	if url, err = getngrokurl.GetngrokURL(path); err != nil {
		showerror.ShowError("Error executing curl. Please verify if ngrok it is up and running.\n")
		os.Exit(1)
	}

	fmt.Println(url)
}
