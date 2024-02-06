package main

import "gotravel/pkg/stdio"

func intro(io stdio.IO) {
	io.Echo("# Bem-vindo(a) ao #gotravel, sua calculadora de custo de viagem (combustível)")
	io.Echo("# Para fazer a calculadora funcionar, responda as perguntas a seguir")
	io.Echo("# ou pressione CTRL+C a qualquer momento para encerrar o programa.")
	io.Echo("# ")
	io.Echo("# Observação:")
	io.Echo("# Para valores numéricos, se quiser informar uma fração, insira somente")
	io.Echo("# números e ponto (.) no lugar da vírgula.")
	io.Echo("# ")
}
