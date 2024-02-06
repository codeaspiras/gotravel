package main

import (
	"gotravel/cmd"
	"gotravel/pkg/stdio"
)

func main() {
	// dependencies
	io := stdio.New()

	// run
	cmd.CLI(io)
}
