package calculator

import (
	"gotravel/pkg/errs"
	"gotravel/pkg/utils"
)

func AskCostPerLiter() (float64, error) {
	v, err := utils.AskFloat("Qual o valor por litro de combustível (R$/L)? > ")
	if err != nil {
		return 0, err
	}

	if v <= 0 {
		return 0, errs.ErrCostPerLiterCannotBeEqualOrLowerThanZero
	}

	return v, nil
}
