package main

import (
	"github.com/codeaspiras/gotravel/cmd"
	"github.com/codeaspiras/gotravel/stdio"
)

func main() {
	// dependencies
	io := stdio.New()

	// run
	cmd.CLI(io)
}
