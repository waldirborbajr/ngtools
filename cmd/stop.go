/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"localhost/ngtools/internal/killprocess"
	"localhost/ngtools/internal/showerror"
	"os"
)

func (cli *Cli) stop() {

	if err := killprocess.KillRunningProcess("ngrok"); err != nil {
		showerror.ShowError(err.Error())
		os.Exit(1)
	}

}
