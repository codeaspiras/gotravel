package calculator

import (
	"gotravel/pkg/constants"
	"gotravel/pkg/errs"
	"gotravel/pkg/stdio"
)

func AskCostPerLiter(io stdio.IO) (float64, error) {
	v, err := io.AskFloat(constants.QuestionCostPerLiter)
	if err != nil {
		return 0, err
	}

	if v <= 0 {
		return 0, errs.ErrCostPerLiterCannotBeEqualOrLowerThanZero
	}

	return v, nil
}
