package main

import (
	"gotravel/calculator"
	"gotravel/pkg/stdio"
	"gotravel/pkg/utils"
)

func main() {
	io := stdio.New()
	defer utils.PanicHandler(io)
	intro(io)

	costPerLiter, err := calculator.AskCostPerLiter(io)
	if err != nil {
		utils.ErrorHandler(io, err)
		return
	}

	distance, err := calculator.AskDistance(io)
	if err != nil {
		utils.ErrorHandler(io, err)
		return
	}

	distancePerLiter, err := calculator.AskDistancePerLiter(io)
	if err != nil {
		utils.ErrorHandler(io, err)
		return
	}

	totalDistance := distance * 2 // go & back
	estimation := calculator.CostToCover(totalDistance, distancePerLiter, costPerLiter)
	io.Echo("O custo estimado para percorrer %.2fkm é de R$%.2f!", totalDistance, estimation)
	io.End()
}
