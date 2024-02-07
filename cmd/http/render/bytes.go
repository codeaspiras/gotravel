package render

import (
	"gotravel/pkg/stdio"
	"net/http"
)

func Bytes(io stdio.IO, w http.ResponseWriter, status int, bytes []byte) {
	w.WriteHeader(status)
	_, err := w.Write(bytes)
	if err != nil {
		echo(io, "could not write bytes: %s", err)
	}
}
