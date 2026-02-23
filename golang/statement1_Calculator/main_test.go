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
		wantErr  bool
	}{
		{"Addition", 2.0, 5.1, "+", 7.1, false},
		{"Subtraction", 3, -5, "-", 8, false},
		{"Multiplication", -6, -2, "*", 12, false},
		{"Division", 7, 3.5, "/", 2, false},
		{"DivisionByZero", 9, 0, "/", 0, true},
		{"InvalidOperator", 8, 8, "9", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator(tt.op1, tt.op2, tt.op)

			if (err != nil) != tt.wantErr {
				t.Errorf("%s: unexpected error status: got error = %v, wantErr = %v", tt.name, err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("%s, got : %v, want : %v", tt.name, got, tt.want)
			}
		})
	}
}
