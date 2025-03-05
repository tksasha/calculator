package calculator

import (
	"math"
	"regexp"

	"github.com/mrxrsd/gojacego"
	"github.com/shopspring/decimal"
)

const ROUND = 2

func Calculate(formula string) (decimal.Decimal, error) {
	if formula == "" {
		return decimal.NewFromInt(0), NewFormulaRequiredError()
	}

	formula = regexp.
		MustCompile("[^0-9+-/*()]").
		ReplaceAllString(formula, "")

	engine, err := gojacego.NewCalculationEngine()
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	result, err := engine.Calculate(formula, nil)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	if result == math.Inf(1) || result == math.Inf(-1) {
		return decimal.NewFromInt(0), NewInfinityError()
	}

	return decimal.NewFromFloat(result).Round(ROUND), nil
}
