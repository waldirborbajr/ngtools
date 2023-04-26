/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package main

import (
	"localhost/ngtools/cmd"
)

func main() {

	// fmt.Println(os.Environ())

	cmd := cmd.Cli{}
	cmd.Run()
}
