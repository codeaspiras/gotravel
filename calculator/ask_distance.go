package calculator

import (
	"gotravel/pkg/errs"
	"gotravel/pkg/utils"
)

func AskDistance() (float64, error) {
	v, err := utils.AskFloat("Qual a distância do ponto de partida até o ponto de chegada (em km)? > ")
	if err != nil {
		return 0, err
	}

	if v <= 0 {
		return 0, errs.ErrDistanceCannotBeEqualOrLowerThanZero
	}

	return v, nil
}
