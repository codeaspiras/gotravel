package render

import (
	"gotravel/pkg/stdio"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidationError(io stdio.IO, w http.ResponseWriter, errors validator.ValidationErrorsTranslations) {
	statusCode := http.StatusBadRequest

	JSON(io, w, statusCode, map[string]interface{}{
		"status": statusCode,
		"errors": errors,
	})
}
