package main

import (
	"testing"
	"math"
)


// TestF2C tests conversion from Fahrenheit to Celsius. 
func TestF2C(t *testing.T) {
	Fval := 33.8
	Cval := F2C(Fval)
	Crequired := 1.0
	// Compare with a tolerance of 10^-10 as we are working with floats.
	if math.Abs(Cval - Crequired) > 1e-10 {
		t.Fatalf("F2C(%v) = %v, want %v", Fval, Cval, Crequired)
	}
}

