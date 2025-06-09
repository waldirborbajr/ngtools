/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
	"localhost/ngtools/internal/getngrokurl"
	"localhost/ngtools/internal/showerror"
	"os"
)

func (cli *Cli) addr() {

	// Get url generated
	url, err := getngrokurl.GetNgrokURL()
	if err != nil {
		showerror.ShowError("Error executing curl. Please verify if ngrok it is up and running.\n")
		os.Exit(1)
	}

	fmt.Println(url)

}
