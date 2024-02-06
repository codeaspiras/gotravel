# GOTRAVEL

Esse projeto foi desenvolvido apenas para resolver um desafio lançado pelo [CodeAspiras](https://blog.codeaspiras.dev).
Caso tenha se interessado pela proposta e queira desenvolver sua solução, sinta-se a vontade!

# PREMISSA

Uma família deseja viajar durante as férias, mas precisam calcular uma previsão de custo com combustível para saber quanto dinheiro devem reservar para tal. As informações que eles podem fornecer ao seu programa são:

A) O valor por litro de combustível (R$/L);
B) A distância da casa deles até o ponto de destino (km);
C) Quantos quilômetros o automóvel consegue percorrer com um litro de combustível (km/L).

# DESAFIO

Faça um programa que pergunte essas três informações e retorne o valor total que a família deve reservar para poder ir e voltar na viagem.

# OBSERVAÇÕES

para simplificar, apenas acredite que a distância para voltar pra casa é a mesma distância da ida (na prática, não é!);
também desconsidere possíveis necessidades de abastecimento durante o trajeto (pois, nesse caso, teriam de saber em quais postos iriam abastecer e quanto custa o litro do combustível lá);
a família inteira vai num automóvel só, mas lembre-se que o custo é de ida E VOLTA!!

# COMO RODAR O PROJETO

1. Você precisar ter o [Golang v1.20+](https://go.dev/) instalado na sua máquina. Caso não o tenha, acesse o link informado anteriormente, baixe e instale!
2. Abra um terminal no diretório raiz desse projeto;
3. Rode `go run .` e siga as instruções do programa!

Você também pode compilar o projeto em um executável rodando `go build .`. Um arquivo executável será criado na raiz do projeto para que possa fazer o que quiser.