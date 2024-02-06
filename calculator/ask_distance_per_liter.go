package calculator

import (
	"gotravel/pkg/errs"
	"gotravel/pkg/utils"
)

func AskDistancePerLiter() (float64, error) {
	v, err := utils.AskFloat("Quantos quilômetros o automóvel consegue percorrer com um litro de combustível (km/L)? > ")
	if err != nil {
		return 0, err
	}

	if v <= 0 {
		return 0, errs.ErrDistancePerLiterCannotBeEqualOrLowerThanZero
	}

	return v, nil
}
