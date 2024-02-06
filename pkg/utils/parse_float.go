package utils

import "strconv"

func ParseFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err.Error())
	}

	return v
}
