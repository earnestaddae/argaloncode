package data

import (
	"errors"
)

// Operands is the struct containing the first and second numbers as operands
type Operands struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

// Divide returns a map of string and float of x/y and y/x
func (op *Operands) Divide(a, b float64) (float64, float64, error) {
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	op.A = a
	op.B = b

	return op.A / op.B, op.B / op.A, nil
}
