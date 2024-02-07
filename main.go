package main

import (
	"gotravel/cmd"
	"gotravel/pkg/stdio"
	"os"
)

const (
	commandServeHTTP = "serve"
	commandRunCLI    = ""
)

func main() {
	// dependencies
	io := stdio.New()

	// run
	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case commandServeHTTP:
		cmd.HTTP(io)
	case commandRunCLI:
		cmd.CLI(io)
	default:
		io.Echo("not implemented")
		io.End()
	}
}
