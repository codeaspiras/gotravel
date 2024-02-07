package render

import (
	"encoding/json"
	"gotravel/pkg/stdio"
	"net/http"
)

func JSON(io stdio.IO, w http.ResponseWriter, status int, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		echo(io, "could not marshal data to render JSON: %s", err)
		bytes = []byte("")
	}

	w.Header().Add("Content-Type", "application/json")
	Bytes(io, w, status, bytes)
}
