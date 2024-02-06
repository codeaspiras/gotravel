package main

import "gotravel/pkg/stdio"

func intro() {
	stdio.Echo("# Bem-vindo(a) ao #gotravel, sua calculadora de custo de viagem (combustível)")
	stdio.Echo("# Para fazer a calculadora funcionar, responda as perguntas a seguir")
	stdio.Echo("# ou pressione CTRL+C a qualquer momento para encerrar o programa.")
	stdio.Echo("# ")
	stdio.Echo("# Observação:")
	stdio.Echo("# Para valores numéricos, se quiser informar uma fração, insira somente")
	stdio.Echo("# números e ponto (.) no lugar da vírgula.")
	stdio.Echo("# ")
}
