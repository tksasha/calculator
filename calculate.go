package calculator

import (
	"go/token"
	"go/types"
	"regexp"
	"strconv"
)

func Calculate(formula string) (float64, error) {
	formula = regexp.
		MustCompile("[^0-9+-/*()]").
		ReplaceAllString(formula, "")

	fset := token.NewFileSet()

	var pos token.Pos

	pkg := types.NewPackage("", "")

	res, err := types.Eval(fset, pkg, pos, formula)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(res.Value.String(), 64)
}
