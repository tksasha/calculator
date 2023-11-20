package formula

import (
	"testing"
)

const M = "\033[31m`%v` was expected, but it is `%v`\033[0m"

func TestCalculateFormula(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		result, _ := Calculate("2+2")

		if result != 4 {
			t.Errorf(`Calculate("2+2") != 4`)
		}
	})

	t.Run("when it contains spaces", func(t *testing.T) {
		result, _ := Calculate(" 2 + 2 ")

		if result != 4 {
			t.Errorf(`Calculate("2+2") != 4`)
		}
	})

	t.Run("when it contains letters", func(t *testing.T) {
		result, _ := Calculate("a2b+c2d")

		if result != 4 {
			t.Errorf(`Calculate("2+2") != 4`)
		}
	})

	t.Run("when it is an addition", func(t *testing.T) {
		result, _ := Calculate("3+4")

		if result != 7 {
			t.Errorf(`Calculate(3+4) != 7`)
		}
	})

	t.Run("when it is a substraction", func(t *testing.T) {
		result, _ := Calculate("8-6")

		if result != 2 {
			t.Errorf(`Calculate(8-6) != 2`)
		}
	})

	t.Run("when it is a multiplication", func(t *testing.T) {
		result, _ := Calculate("4*5")

		if result != 20 {
			t.Errorf(`Calculate(4*5) != 20`)
		}
	})

	t.Run("when it is a division", func(t *testing.T) {
		result, _ := Calculate("9/3")

		if result != 3 {
			t.Errorf(`Calculate(9/3) != 3`)
		}
	})

	t.Run("when it is empty", func(t *testing.T) {
		result, _ := Calculate("")

		if result != 0 {
			t.Errorf(`Calculate("") != 0`)
		}
	})

	t.Run("when it contains brackets", func(t *testing.T) {
		result, _ := Calculate("(5 + 3) * 4")

		if result != 32 {
			t.Errorf(`Calculate((5 + 3) * 4) != 32`)
		}
	})

	t.Run("when calculated result is float64", func(t *testing.T) {
		exp := 111.11

		res, _ := Calculate("42.1 + 69.01")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})
}
