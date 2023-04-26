/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"fmt"
)

// type startCmd struct{}

func (cli *Cli) start(protocol string, port int) {

	fmt.Println("started... ", protocol, port)

}
