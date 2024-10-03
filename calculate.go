package calculator

import (
	"math"
	"regexp"

	"github.com/mrxrsd/gojacego"
)

func Calculate(formula string) (float64, error) {
	if formula == "" {
		return 0, NewFormulaRequiredError()
	}

	formula = regexp.
		MustCompile("[^0-9+-/*()]").
		ReplaceAllString(formula, "")

	engine, err := gojacego.NewCalculationEngine()
	if err != nil {
		return 0, err
	}

	result, err := engine.Calculate(formula, nil)
	if err != nil {
		return 0, err
	}

	if result == math.Inf(1) || result == math.Inf(-1) {
		return 0, NewInfinityError()
	}

	return math.Round(result*100) / 100, nil //nolint:mnd
}
