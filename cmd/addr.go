/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
	"localhost/ngtools/internal/getngrokurl"
	"localhost/ngtools/internal/hascurl"
	"localhost/ngtools/internal/showerror"
	"os"
)

func (cli *Cli) addr() {

	path, err := hascurl.HasCurl()
	if err != nil {
		showerror.ShowError("Curl not found. Please install it first and run it again.\n")
		os.Exit(1)
	}

	// Get url generated
	url, err := getngrokurl.GetngrokURL(path)
	if err != nil {
		showerror.ShowError("Error executing curl. Please verify if ngrok it is up and running.\n")
		os.Exit(1)
	}

	fmt.Println(url)

}
