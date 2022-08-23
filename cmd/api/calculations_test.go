package main

import (
	"testing"
)

// import "testing"

func TestDivision(t *testing.T) {
	tests := []struct {
		name         string
		a, b         float64
		wanta, wantb float64
	}{
		{name: "a > b", a: 100, b: 2, wanta: 100.0 / 2.0, wantb: 2.0 / 100.0},
		{name: "a = b", a: 1, b: 1, wanta: 1, wantb: 1},
		{name: "a < b", a: 3, b: 9, wanta: 3.0 / 9.0, wantb: 9.0 / 3.0},
		{name: "b = 0", a: 8, b: 0, wanta: 0, wantb: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b, _ := testApp.Division(tt.a, tt.b)
			if a != tt.wanta {
				t.Errorf("got %f ; want %f", a, tt.wanta)
			}
			if b != tt.wantb {
				t.Errorf("got %f ; want %f", b, tt.wantb)
			}
		})
	}

}
