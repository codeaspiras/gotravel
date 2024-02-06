package main

import (
	"gotravel/calculator"
	"gotravel/pkg/stdio"
	"gotravel/pkg/utils"
)

func main() {
	defer utils.PanicHandler()
	intro()

	costPerLiter, err := calculator.AskCostPerLiter()
	if err != nil {
		utils.ErrorHandler(err)
		return
	}

	distance, err := calculator.AskDistance()
	if err != nil {
		utils.ErrorHandler(err)
		return
	}

	distancePerLiter, err := calculator.AskDistancePerLiter()
	if err != nil {
		utils.ErrorHandler(err)
		return
	}

	totalDistance := distance * 2 // go & back
	estimation := calculator.CostToCover(totalDistance, distancePerLiter, costPerLiter)
	stdio.Echo("O custo estimado para percorrer %.2fkm é de R$%.2f!", totalDistance, estimation)
	utils.End()
}
