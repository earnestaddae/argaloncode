package main

import "errors"

// Divide returns a map of string and float of x/y and y/x
func (app *application) Division(a, b float64) (float64, float64, error) {
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}

	return a / b, b / a, nil
}
