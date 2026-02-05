package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name     string
		op1, op2 float64
		op       string
		want     float64
	}{
		{"Addition", 2.0, 5.1, "+", 7.1},
		{"Subtraction", 3, -5, "-", 8},
		{"Multiplication", -6, -2, "*", 12},
		{"Division", 7, 3.5, "/", 2},
		{"DivisionByZero", 9, 0, "/", 0},
		{"InvalidOperator", 8, 8, "9", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator(tt.op1, tt.op2, tt.op)

			if got != tt.want {
				t.Errorf("%s, got : %v, want : %v", tt.name, got, tt.want)
			}
		})
	}
}
