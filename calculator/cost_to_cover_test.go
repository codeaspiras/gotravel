package calculator_test

import (
	"gotravel/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

type costToCover_TestCase struct {
	distancePerLiter float64
	distance         float64
	costPerLiter     float64
	expectedResult   float64
}

func costToCover_TestCases() map[string]costToCover_TestCase {
	return map[string]costToCover_TestCase{
		"should result $53.50": costToCover_TestCase{
			distancePerLiter: 10,
			distance:         100,
			costPerLiter:     5.35,
			expectedResult:   53.5,
		},
		"should result $72.50": costToCover_TestCase{
			distancePerLiter: 20,
			distance:         500,
			costPerLiter:     2.9,
			expectedResult:   72.5,
		},
	}
}

func Test_CostToCover(t *testing.T) {
	cases := costToCover_TestCases()
	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			// Arrange
			ctc := calculator.CostToCover(c.distance, c.distancePerLiter, c.costPerLiter)

			// Assert
			assert.Equal(t, c.expectedResult, ctc)
		})
	}
}
