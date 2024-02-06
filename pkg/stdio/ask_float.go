package stdio

import (
	"strconv"
)

func (i *io) AskFloat(template string, args ...interface{}) (float64, error) {
	v, err := i.Ask(template, args...)
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}
