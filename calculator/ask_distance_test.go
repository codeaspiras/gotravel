package calculator_test

import (
	"errors"
	"gotravel/calculator"
	"gotravel/mocks"
	"gotravel/pkg/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AskDistance(t *testing.T) {
	io := new(mocks.IO)
	question := "Qual a distância do ponto de partida até o ponto de chegada (em km)? > "

	t.Run("should return the inputed number", func(t *testing.T) {
		// Arrange
		expected := float64(123.5)
		io.On("AskFloat", question).Return(expected, nil).Once()

		// Act
		v, err := calculator.AskDistance(io)

		// Assert
		assert.NoError(t, err, "err should be nil")
		assert.Equal(t, expected, v, "v should be equal to the expected value")
	})

	t.Run("should not return zero", func(t *testing.T) {
		// Arrange
		io.On("AskFloat", question).Return(float64(0), nil).Once()

		// Act
		_, err := calculator.AskDistance(io)

		// Assert
		assert.Equal(t, errs.ErrDistanceCannotBeEqualOrLowerThanZero, err, "err should be errs.ErrDistanceCannotBeEqualOrLowerThanZero")
	})

	t.Run("should not return numbers lower than zero", func(t *testing.T) {
		// Arrange
		io.On("AskFloat", question).Return(float64(-12), nil).Once()

		// Act
		_, err := calculator.AskDistance(io)

		// Assert
		assert.Equal(t, errs.ErrDistanceCannotBeEqualOrLowerThanZero, err, "err should be errs.ErrDistanceCannotBeEqualOrLowerThanZero")
	})

	t.Run("should return error", func(t *testing.T) {
		// Arrange
		expected := errors.New("testing error")
		io.On("AskFloat", question).Return(float64(0), expected).Once()

		// Act
		_, err := calculator.AskDistance(io)

		// Assert
		assert.Equal(t, expected, err, "err should be equal to the expected")
	})
}
