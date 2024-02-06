package cmd

import (
	"gotravel/calculator"
	"gotravel/pkg/stdio"
	"gotravel/pkg/utils"
)

func CLI(io stdio.IO) {
	defer utils.PanicHandler(io)

	io.Echo("# Bem-vindo(a) ao #gotravel, sua calculadora de custo de viagem (combustível)")
	io.Echo("# Para fazer a calculadora funcionar, responda as perguntas a seguir")
	io.Echo("# ou pressione CTRL+C a qualquer momento para encerrar o programa.")
	io.Echo("# ")
	io.Echo("# Observação:")
	io.Echo("# Para valores numéricos, se quiser informar uma fração, insira somente")
	io.Echo("# números e ponto (.) no lugar da vírgula.")
	io.Echo("# ")

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
