package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrEmptyExpression   = errors.New("empty expression")
	ErrDivisionByZero    = errors.New("division by zero")
	ErrNotCorrInput      = errors.New("not correct input")
)
