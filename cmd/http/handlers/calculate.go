package handlers

import (
	"encoding/json"
	"fmt"
	"gotravel/calculator"
	"gotravel/cmd/http/render"
	"gotravel/pkg/altcontext"
	ioreader "io"
	"net/http"
)

type CalculateInput struct {
	CostPerLiter     float64 `json:"cost_per_liter" validate:"required,gt=0"`
	Distance         float64 `json:"distance" validate:"required,gt=0"`
	DistancePerLiter float64 `json:"distance_per_liter" validate:"required,gt=0"`
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	io, err := altcontext.GetIO(r.Context())
	if err != nil {
		render.Error(nil, w, err, http.StatusInternalServerError)
		return
	}

	validator, err := altcontext.GetValidator(r.Context())
	if err != nil {
		render.Error(io, w, err, http.StatusInternalServerError)
		return
	}

	bodyBytes, err := ioreader.ReadAll(r.Body)
	if err != nil {
		render.Error(io, w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var input CalculateInput
	err = json.Unmarshal(bodyBytes, &input)
	if err != nil {
		render.Error(io, w, err, http.StatusBadRequest)
		return
	}

	errors := validator.Validate(input)
	if len(errors) > 0 {
		render.ValidationError(io, w, errors)
	}

	totalDistance := input.Distance * 2
	result := calculator.CostToCover(
		totalDistance,
		input.DistancePerLiter,
		input.CostPerLiter,
	)

	render.JSON(io, w, http.StatusOK, map[string]interface{}{
		"totalDistance": fmt.Sprintf("%.2f km", totalDistance),
		"cost":          fmt.Sprintf("$%.2f", result),
	})
}
