/*
Copyright © 2023 Waldir Borba Junior <wborbajr@gmail.com>
*/

package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct{}

func (cli *Cli) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("   start [protocol] [port] - Start ngrok PROTOCOL to be used http/tcp/tls PORT to be opened")
	fmt.Println("   stop - Stop ngrok")
}

func (cli *Cli) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *Cli) Run() {
	cli.validateArgs()

	startCmd := flag.NewFlagSet("start", flag.ExitOnError)

	startProto := startCmd.String("protocol", "", "The protocolo to be enabled http/tls/tcp")
	startPort := startCmd.Int("port", 0, "The port to be opened")

	switch os.Args[1] {
	case "start":
		if err := startCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if startCmd.Parsed() {
		if *startProto == "" || *startPort <= 0 {
			startCmd.Usage()
			os.Exit(1)
		}

		cli.start(*startProto, *startPort)
	}
}
