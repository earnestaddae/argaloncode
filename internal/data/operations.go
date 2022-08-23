package data

import (
	"fmt"
)

// Operands is the struct containing the first and second numbers as operands
type Operands struct {
	FirstNumber  float64 `json:"first_number"`
	SecondNumber float64 `json:"second_number"`
}

// Divide returns a map of string and float of x/y and y/x
func (op *Operands) Divide(x, y float64) (map[string]float64, error) {
	if y == 0 {
		return nil, fmt.Errorf("can't divide %0.4f by zero", x)
	}

	op.FirstNumber = x
	op.SecondNumber = y

	result := map[string]float64{
		"first_number / second_number": op.FirstNumber / op.SecondNumber,
		"second_number / first_number": op.SecondNumber / op.FirstNumber,
	}

	return result, nil
}
