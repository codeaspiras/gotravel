package utils

import (
	"gotravel/pkg/stdio"
	"strconv"
)

func AskFloat(template string, args ...interface{}) (float64, error) {
	v, err := stdio.Ask(template, args...)
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}
