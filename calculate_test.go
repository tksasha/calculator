package calculator_test

import (
	"testing"

	"github.com/tksasha/calculator"
	"gotest.tools/v3/assert"
)

func TestCalculateFormula(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		res, _ := calculator.Calculate("2+2")

		assert.Equal(t, res, 4.0)
	})

	t.Run("when it contains spaces", func(t *testing.T) {
		res, _ := calculator.Calculate(" 2 + 2 ")

		assert.Equal(t, res, 4.0)
	})

	t.Run("when it contains letters", func(t *testing.T) {
		res, _ := calculator.Calculate("a2b+c2d")

		assert.Equal(t, res, 4.0)
	})

	t.Run("when it is an addition", func(t *testing.T) {
		res, _ := calculator.Calculate("3+4")

		assert.Equal(t, res, 7.0)
	})

	t.Run("when it is a subtraction", func(t *testing.T) {
		res, _ := calculator.Calculate("8-6")

		assert.Equal(t, res, 2.0)
	})

	t.Run("when it is a multiplication", func(t *testing.T) {
		res, _ := calculator.Calculate("4*5")

		assert.Equal(t, res, 20.0)
	})

	t.Run("when it is a division", func(t *testing.T) {
		res, _ := calculator.Calculate("9/3")

		assert.Equal(t, res, 3.0)
	})

	t.Run("when it is empty", func(t *testing.T) {
		res, _ := calculator.Calculate("")

		assert.Equal(t, res, 0.0)
	})

	t.Run("when it contains brackets", func(t *testing.T) {
		res, _ := calculator.Calculate("(5 + 3) * 4")

		assert.Equal(t, res, 32.0)
	})

	t.Run("when calculated result is float64", func(t *testing.T) {
		res, _ := calculator.Calculate("42.1 + 69.01")

		assert.Equal(t, res, 111.11)
	})

	t.Run("when try to calculate 4269.69+6942.69 it should return 11212.11", func(t *testing.T) {
		res, _ := calculator.Calculate("4269.42+6942.69")

		assert.Equal(t, res, 11_212.11)
	})
}
