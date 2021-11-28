package formula

import (
	"testing"
)

func TestCalculateWhenFormulaIsValid(t *testing.T) {
	result, _ := Calculate("2+2")

	if result != 4 {
		t.Errorf(`Calculate("2+2") != 4`)
	}
}

func TestCalculateWhenFormulaContainsSpaces(t *testing.T) {
	result, _ := Calculate(" 2 + 2 ")

	if result != 4 {
		t.Errorf(`Calculate("2+2") != 4`)
	}
}

func TestCalculateWhenFormulaContainsLetters(t *testing.T) {
	result, _ := Calculate("a2b+c2d")

	if result != 4 {
		t.Errorf(`Calculate("2+2") != 4`)
	}
}

func TestCalculateAddition(t *testing.T) {
	result, _ := Calculate("3+4")

	if result != 7 {
		t.Errorf(`Calculate(3+4) != 7`)
	}
}

func TestCalculateSubtraction(t *testing.T) {
	result, _ := Calculate("8-6")

	if result != 2 {
		t.Errorf(`Calculate(8-6) != 2`)
	}
}

func TestCalculateMultiplication(t *testing.T) {
	result, _ := Calculate("4*5")

	if result != 20 {
		t.Errorf(`Calculate(4*5) != 20`)
	}
}

func TestCalculateDivision(t *testing.T) {
	result, _ := Calculate("9/3")

	if result != 3 {
		t.Errorf(`Calculate(9/3) != 3`)
	}
}

func TestCalculateWhenFormulaIsEmpty(t *testing.T) {
	result, _ := Calculate("")

	if result != 0 {
		t.Errorf(`Calculate("") != 0`)
	}
}
