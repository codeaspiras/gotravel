package utils

import "github.com/codeaspiras/gotravel/stdio"

func PanicHandler(io stdio.IO) {
	err := recover()
	if err == nil {
		return
	}

	io.Echo("# Isso gerou um erro fatal:\n- %s", err)
	io.End()
}
