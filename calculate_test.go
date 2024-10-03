package calculator_test

import (
	"testing"

	"github.com/tksasha/calculator"
	"gotest.tools/v3/assert"
)

func TestCalculateFormula(t *testing.T) { //nolint:funlen
	t.Run("when it is valid", func(t *testing.T) {
		res, err := calculator.Calculate("2+2")

		assert.Equal(t, res, 4.0)
		assert.NilError(t, err)
	})

	t.Run("when it contains spaces it should trim them", func(t *testing.T) {
		res, err := calculator.Calculate(" 2 + 2 ")

		assert.Equal(t, res, 4.0)
		assert.NilError(t, err)
	})

	t.Run("when it contains letters it should trim them", func(t *testing.T) {
		res, err := calculator.Calculate("a2b+c2d")

		assert.Equal(t, res, 4.0)
		assert.NilError(t, err)
	})

	t.Run("when it is an addition it should calculate it", func(t *testing.T) {
		res, err := calculator.Calculate("3+4")

		assert.Equal(t, res, 7.0)
		assert.NilError(t, err)
	})

	t.Run("when it is a subtraction it should calculate it", func(t *testing.T) {
		res, err := calculator.Calculate("8-6")

		assert.Equal(t, res, 2.0)
		assert.NilError(t, err)
	})

	t.Run("when it is a multiplication it should calculate it", func(t *testing.T) {
		res, err := calculator.Calculate("4*5")

		assert.Equal(t, res, 20.0)
		assert.NilError(t, err)
	})

	t.Run("when it is a division it should calculate it", func(t *testing.T) {
		res, err := calculator.Calculate("9/3")

		assert.Equal(t, res, 3.0)
		assert.NilError(t, err)
	})

	t.Run("when it is empty it should return an error", func(t *testing.T) {
		res, err := calculator.Calculate("")

		assert.Equal(t, res, 0.0)
		assert.Error(t, err, "formula is required")
	})

	t.Run("when it contains brackets it should calculate it", func(t *testing.T) {
		res, err := calculator.Calculate("(5 + 3) * 4")

		assert.Equal(t, res, 32.0)
		assert.NilError(t, err)
	})

	t.Run("when calculated result is float64 it should round it", func(t *testing.T) {
		res, err := calculator.Calculate("42.111 + 69.014")

		assert.Equal(t, res, 111.13)
		assert.NilError(t, err)
	})

	t.Run("when try to calculate 4269.69+6942.69 it should return 11212.11", func(t *testing.T) {
		res, _ := calculator.Calculate("4269.42+6942.69")

		assert.Equal(t, res, 11_212.11)
	})

	t.Run("when try to calculate 2/0 it should return an error", func(t *testing.T) {
		res, err := calculator.Calculate("2/0")

		assert.Equal(t, res, 0.0)
		assert.Error(t, err, "result is an infinity")
	})
}
