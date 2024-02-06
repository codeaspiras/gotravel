package calculator

import (
	"gotravel/pkg/constants"
	"gotravel/pkg/errs"
	"gotravel/pkg/stdio"
)

func AskDistancePerLiter(io stdio.IO) (float64, error) {
	v, err := io.AskFloat(constants.QuestionDistancePerLiter)
	if err != nil {
		return 0, err
	}

	if v <= 0 {
		return 0, errs.ErrDistancePerLiterCannotBeEqualOrLowerThanZero
	}

	return v, nil
}
