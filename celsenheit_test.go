package main

import (
	"testing"
	"math"
)


func TestF2C(t *testing.T) {
	Fval := 33.8
	Cval := F2C(Fval)
	Crequired := 1.0
	// Compare with a tolerance of 10^-10 as we are working with floats.
	if math.Abs(Cval - Crequired) > 1e-10 {
		t.Fatalf("F2C(%v) = %v, expected %v", Fval, Cval, Crequired)
	}
}


func TestC2F(t *testing.T) {
	Cval := 0.0
	Fval := C2F(Cval)
	Frequired := 32.0
	// Compare with a tolerance of 10^-10 as we are working with floats.
	if math.Abs(Fval - Frequired) > 1e-10 {
		t.Fatalf("F2C(%v) = %v, expected %v", Cval, Fval, Frequired)
	}
}


func TestRandFloat(t *testing.T) {
	min, max := 7.0, 11.0
	val := RandFloat(min, max)
	if (val < min) || (val > max) {
		t.Fatalf("Expected a value between %v and %v, got %v", min, max, val)
	}
}


func TestContainsStr(t *testing.T) {
	slice := []string{"foo", "bar", "baz"}
	// Test with a string that is in the slice.
	containsFoo := ContainsStr(slice, "foo")
	if containsFoo != true {
		t.Fatalf("ContainsStr(slice, 'foo') = false, expected true")
	}
	// Test with a string that is not in the slice.
	containsQux := ContainsStr(slice, "qux")
	if containsQux != false {
		t.Fatalf("ContainsStr(slice, 'qux') = true, expected false")
	}
}


func TestVerboseDegreeConversion(t *testing.T) {
	// Test for a valid C -> F conversion.
	msg, err := VerboseDegreeConversion(32.0, "C", "F")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	expectedMsg := "32°C is equivalent to 89.6°F.\n"
	if msg != expectedMsg {
		t.Fatalf("VerboseDegreeConversion(32.0, 'C', 'F') gives \n %s expected \n %s",
		msg, expectedMsg)
	}

	// Test for a valid C -> F conversion.
	msg, err = VerboseDegreeConversion(100.0, "F", "C")
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	expectedMsg = "100°F is equivalent to 37.8°C.\n"
	if msg != expectedMsg {
		t.Fatalf("VerboseDegreeConversion(100.0, 'F', 'C') gives \n %s expected \n %s",
		msg, expectedMsg)
	}

	// Test in the case of an invalid input scale pair.
	_, err = VerboseDegreeConversion(1.0, "A", "B")
	if err == nil {
		t.Fatal("Expected error upon invalid degree conversion.")
	}
}


func TestRunGuess(t *testing.T) {
}


func TestJudgeGuess(t *testing.T) {
}


func TestMain(t *testing.T) {
}

