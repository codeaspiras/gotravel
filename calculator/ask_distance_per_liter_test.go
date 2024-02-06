package calculator_test

import (
	"errors"
	"gotravel/calculator"
	"gotravel/mocks"
	"gotravel/pkg/constants"
	"gotravel/pkg/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AskDistancePerLiter(t *testing.T) {
	io := new(mocks.IO)

	t.Run("should return the inputed number", func(t *testing.T) {
		// Arrange
		expected := float64(123.5)
		io.On("AskFloat", constants.QuestionDistancePerLiter).Return(expected, nil).Once()

		// Act
		v, err := calculator.AskDistancePerLiter(io)

		// Assert
		assert.NoError(t, err, "err should be nil")
		assert.Equal(t, expected, v, "v should be equal to the expected value")
		assert.True(t, io.AssertExpectations(t), "io expectations should be true")
	})

	t.Run("should not return zero", func(t *testing.T) {
		// Arrange
		io.On("AskFloat", constants.QuestionDistancePerLiter).Return(float64(0), nil).Once()

		// Act
		_, err := calculator.AskDistancePerLiter(io)

		// Assert
		assert.Equal(t, errs.ErrDistancePerLiterCannotBeEqualOrLowerThanZero, err, "err should be errs.ErrDistancePerLiterCannotBeEqualOrLowerThanZero")
		assert.True(t, io.AssertExpectations(t), "io expectations should be true")
	})

	t.Run("should not return numbers lower than zero", func(t *testing.T) {
		// Arrange
		io.On("AskFloat", constants.QuestionDistancePerLiter).Return(float64(-12), nil).Once()

		// Act
		_, err := calculator.AskDistancePerLiter(io)

		// Assert
		assert.Equal(t, errs.ErrDistancePerLiterCannotBeEqualOrLowerThanZero, err, "err should be errs.ErrDistancePerLiterCannotBeEqualOrLowerThanZero")
		assert.True(t, io.AssertExpectations(t), "io expectations should be true")
	})

	t.Run("should return error", func(t *testing.T) {
		// Arrange
		expected := errors.New("testing error")
		io.On("AskFloat", constants.QuestionDistancePerLiter).Return(float64(0), expected).Once()

		// Act
		_, err := calculator.AskDistancePerLiter(io)

		// Assert
		assert.Equal(t, expected, err, "err should be equal to the expected")
		assert.True(t, io.AssertExpectations(t), "io expectations should be true")
	})
}
