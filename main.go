package main

import (
	"gotravel/calculator"
	"gotravel/pkg/stdio"
	"gotravel/pkg/utils"
)

func main() {
	defer utils.PanicHandler()
	intro()

	// R$/L
	costPerLiter := utils.ParseFloat(
		stdio.Ask("Qual o valor por litro de combustível (R$/L)? > "),
	)
	if costPerLiter <= 0 {
		panic("(price <= 0)\nSe está de graça, essa calculadora não tem serventia alguma...")
	}

	// KM
	distance := utils.ParseFloat(
		stdio.Ask("Qual a distância do ponto de partida até o ponto de chegada (em km)? > "),
	)
	if distance <= 0 {
		panic("(distance <= 0)\nOra, então você não vai ter custo. Pra quê fazer conta?")
	}

	// KM/L
	distancePerLiter := utils.ParseFloat(
		stdio.Ask("Quantos quilômetros o automóvel consegue percorrer com um litro de combustível (km/L)? > "),
	)

	if distancePerLiter <= 0 {
		panic("(distancePerLiter <= 0)\nIsso é impossível. Seu automóvel é do futuro?")
	}

	totalDistance := distance * 2 // go & back
	estimation := calculator.CostToCover(totalDistance, distancePerLiter, costPerLiter)
	stdio.Echo("O custo estimado para percorrer %.2fkm é de R$%.2f!", totalDistance, estimation)
	utils.End()
}
