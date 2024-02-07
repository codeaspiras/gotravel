package render

import (
	"gotravel/pkg/errs"
	"gotravel/pkg/stdio"
	"net/http"
)

func Error(io stdio.IO, w http.ResponseWriter, err error, statuses ...int) {
	statusCode := http.StatusInternalServerError
	if len(statuses) > 0 {
		statusCode = statuses[0]
	} else {
		switch err {
		case errs.ErrCostPerLiterCannotBeEqualOrLowerThanZero, errs.ErrDistanceCannotBeEqualOrLowerThanZero, errs.ErrDistancePerLiterCannotBeEqualOrLowerThanZero:
			statusCode = http.StatusBadRequest
		}
	}

	JSON(io, w, statusCode, map[string]interface{}{
		"status": statusCode,
		"error":  err.Error(),
	})
}
