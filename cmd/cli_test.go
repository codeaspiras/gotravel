package cmd_test

import (
	"errors"
	"gotravel/cmd"
	"gotravel/mocks"
	"gotravel/pkg/constants"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func expectIntroduction(io *mocks.IO) {
	io.On("Echo", "# Bem-vindo(a) ao #gotravel, sua calculadora de custo de viagem (combustível)").Once()
	io.On("Echo", "# Para fazer a calculadora funcionar, responda as perguntas a seguir").Once()
	io.On("Echo", "# ou pressione CTRL+C a qualquer momento para encerrar o programa.").Once()
	io.On("Echo", "# ").Once()
	io.On("Echo", "# Observação:").Once()
	io.On("Echo", "# Para valores numéricos, se quiser informar uma fração, insira somente").Once()
	io.On("Echo", "# números e ponto (.) no lugar da vírgula.").Once()
	io.On("Echo", "# ").Once()
}

func Test_CLI(t *testing.T) {
	io := new(mocks.IO)

	t.Run("should return the calculation perfectly", func(t *testing.T) {
		// Arrange
		costPerLiter := float64(5.35)
		distance := float64(50)
		distancePerLiter := float64(10)
		expectedResult := float64(53.5)

		// Mocks
		expectIntroduction(io)
		io.On("AskFloat", constants.QuestionCostPerLiter).Return(costPerLiter, nil).Once()
		io.On("AskFloat", constants.QuestionDistance).Return(distance, nil).Once()
		io.On("AskFloat", constants.QuestionDistancePerLiter).Return(distancePerLiter, nil).Once()
		io.On("Echo", "O custo estimado para percorrer %.2fkm é de R$%.2f!", float64(100), expectedResult).Once()
		io.On("End").Once()

		// Act
		cmd.CLI(io)

		// Assert
		assert.True(t, io.AssertExpectations(t), "io expectations should be true")
	})

	t.Run("should echo error on distance per liter question fails", func(t *testing.T) {
		// Arrange
		costPerLiter := float64(5.35)
		distance := float64(50)
		testError := errors.New("test error")

		// Mocks
		expectIntroduction(io)
		io.On("AskFloat", constants.QuestionCostPerLiter).Return(costPerLiter, nil).Once()
		io.On("AskFloat", constants.QuestionDistance).Return(distance, nil).Once()
		io.On("AskFloat", constants.QuestionDistancePerLiter).Return(float64(0), testError).Once()
		io.On("Echo", "# Isso gerou um erro:\n- %s", testError)
		io.On("End").Once()

		// Act
		cmd.CLI(io)

		// Assert
		assert.True(t, io.AssertExpectations(t), "io expectations should be true")
	})

	t.Run("should echo error on distance question fails", func(t *testing.T) {
		// Arrange
		costPerLiter := float64(5.35)
		testError := errors.New("test error")

		// Mocks
		expectIntroduction(io)
		io.On("AskFloat", constants.QuestionCostPerLiter).Return(costPerLiter, nil).Once()
		io.On("AskFloat", constants.QuestionDistance).Return(float64(0), testError).Once()
		io.On("Echo", "# Isso gerou um erro:\n- %s", testError)
		io.On("End").Once()

		// Act
		cmd.CLI(io)

		// Assert
		assert.True(t, io.AssertExpectations(t), "io expectations should be true")
	})

	t.Run("should echo error on cost per liter question fails", func(t *testing.T) {
		// Arrange
		testError := errors.New("test error")

		// Mocks
		expectIntroduction(io)
		io.On("AskFloat", constants.QuestionCostPerLiter).Return(float64(0), testError).Once()
		io.On("Echo", "# Isso gerou um erro:\n- %s", testError)
		io.On("End").Once()

		// Act
		cmd.CLI(io)

		// Assert
		assert.True(t, io.AssertExpectations(t), "io expectations should be true")
	})

	t.Run("should echo panic handler after panic", func(t *testing.T) {
		// Arrange
		testError := errors.New("test error")

		// Mocks
		expectIntroduction(io)
		io.On("AskFloat", constants.QuestionCostPerLiter).Run(func(args mock.Arguments) {
			panic(testError)
		}).Once()
		io.On("Echo", "# Isso gerou um erro fatal:\n- %s", testError)
		io.On("End").Once()

		// Act
		cmd.CLI(io)

		// Assert
		assert.True(t, io.AssertExpectations(t), "io expectations should be true")
	})
}
