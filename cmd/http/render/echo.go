package render

import (
	"gotravel/pkg/stdio"
)

func echo(io stdio.IO, template string, args ...interface{}) {
	if io == nil {
		io = stdio.New()
	}

	io.Echo(template, args...)
}
