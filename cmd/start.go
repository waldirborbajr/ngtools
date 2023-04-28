/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
	"localhost/ngtools/internal/hascurl"
	"localhost/ngtools/internal/showerror"
	"localhost/ngtools/internal/startngrok"
	"os"
)

// type startCmd struct{}

func (cli *Cli) start(protocol string, port int) {
	var path string
	var err error

	if path, err = hascurl.HasCurl(); err != nil {
		showerror.ShowError("Curl not found. Please install it first and run it again.\n")
		os.Exit(1)
	}

	fmt.Println(path)
	startngrok.StartNGRok("http", "8080")

}
