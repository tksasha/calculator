package calculator

type InfinityError struct {
	message string
}

func NewInfinityError() error {
	return &InfinityError{
		message: "result is an infinity",
	}
}

func (e *InfinityError) Error() string {
	return e.message
}

type FormulaRequiredError struct {
	message string
}

func NewFormulaRequiredError() error {
	return &FormulaRequiredError{
		message: "formula is required",
	}
}

func (e *FormulaRequiredError) Error() string {
	return e.message
}
