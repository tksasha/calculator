package formula

import (
	"testing"
)

const M = "\033[31m`%v` was expected, but it is `%v`\033[0m"

func TestCalculateFormula(t *testing.T) {
	t.Run("when it is valid", func(t *testing.T) {
		exp := 4.0

		res, _ := Calculate("2+2")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it contains spaces", func(t *testing.T) {
		exp := 4.0

		res, _ := Calculate(" 2 + 2 ")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it contains letters", func(t *testing.T) {
		exp := 4.0

		res, _ := Calculate("a2b+c2d")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is an addition", func(t *testing.T) {
		exp := 7.0

		res, _ := Calculate("3+4")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is a substraction", func(t *testing.T) {
		exp := 2.0

		res, _ := Calculate("8-6")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is a multiplication", func(t *testing.T) {
		exp := 20.0

		res, _ := Calculate("4*5")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is a division", func(t *testing.T) {
		exp := 3.0

		res, _ := Calculate("9/3")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is empty", func(t *testing.T) {
		exp := 0.0

		res, _ := Calculate("")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it contains brackets", func(t *testing.T) {
		exp := 32.0

		res, _ := Calculate("(5 + 3) * 4")

		if exp != res {
			t.Errorf(M, exp, res)
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
