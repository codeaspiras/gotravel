package utils

import "gotravel/pkg/stdio"

func PanicHandler() {
	err := recover()
	if err == nil {
		return
	}

	stdio.Echo("# Isso gerou um erro fatal:\n- %s", err)
	End()
}
