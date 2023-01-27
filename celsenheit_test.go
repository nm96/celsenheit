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
		t.Fatalf("F2C(%v) = %v, expected %v", Fval, Cval, Crequired)
	}
}
// (Testing F2C explicitly as above is probably not necessary, as it will be
// covered when testing some of the larger functions.)


// TestContainsStr tests the homemade slice contains checker containsStr()
func TestContainsStr(t *testing.T) {
	slice := []string{"foo", "bar", "baz"}
	// Test with a string that is in the slice.
	containsFoo := containsStr(slice, "foo")
	if containsFoo != true {
		t.Fatalf("containsStr(slice, 'foo') = false, expected true")
	}
	// Test with a string that is not in the slice.
	containsQux := containsStr(slice, "qux")
	if containsQux != false {
		t.Fatalf("containsStr(slice, 'qux') = true, expected false")
	}
}
