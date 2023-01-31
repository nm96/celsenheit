package main

import (
	"testing"
	"math"
	"math/rand"
	"strings"
	"bufio"
	"fmt"
)


// floatEqual compares if two floating point values are equal up to some
// tolerance value.
func floatEqual(a, b, tol float64) bool {
	return math.Abs(a - b) < tol
}



func TestF2C(t *testing.T) {
	Fval := 33.8
	Cval := F2C(Fval)
	Crequired := 1.0
	// Compare with a tolerance of 10^-10 as we are working with floats.
	if !floatEqual(Cval, Crequired, 1e-10) {
		t.Fatalf("F2C(%v) = %v, expected %v", Fval, Cval, Crequired)
	}
}


func TestC2F(t *testing.T) {
	Cval := 0.0
	Fval := C2F(Cval)
	Frequired := 32.0
	// Compare with a tolerance of 10^-10 as we are working with floats.
	if !floatEqual(Fval, Frequired, 1e-10) {
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

	// Test for a valid F -> C conversion.
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


func TestNewQuestion(t *testing.T) {
	Cmin, Cmax := -50.0, 50.0
	Fmin, Fmax := C2F(Cmin), C2F(Cmax)
	scales := []string{"C", "F"}
	for i := int64(0); i < 5; i++ {
		rand.Seed(i)
		Q0 := NewQuestion(Cmin, Cmax)
		if !ContainsStr(scales, Q0.fromScale) {
			t.Fatal("Invalid fromScale, expected 'C' or 'F'.")
		}
		if !ContainsStr(scales, Q0.toScale) {
			t.Fatal("Invalid toScale, expected 'C' or 'F'.")
		}
		if Q0.fromScale == "C" {
			if (Q0.val < Cmin) || (Q0.val > Cmax) {
				t.Fatal("Question value outside expected range.")
			}
			if !floatEqual(Q0.ans, C2F(Q0.val), 1e-10) {
				t.Fatal("Question has incorrect ans value.")
			}
		} else {
			if (Q0.val < Fmin) || (Q0.val > Fmax) {
				t.Fatal("Question value outside expected range.")
			}
			if !floatEqual(Q0.ans, F2C(Q0.val), 1e-10) {
				t.Fatal("Question has incorrect ans value.")
			}
		}
		if Q0.guess != 0 {
			t.Fatal("Guess value not initialized to zero.")
		}
	}
}


func TestGetGuess(t *testing.T) {
	// Test with valid guess
	Q := NewQuestion(-50.0, 50.0)
	guessStr := "10.0\n"
	reader := bufio.NewReader(strings.NewReader(guessStr))
	// TODO: Prevent GetGuess from outputtng to consol while testing.
	err := GetGuess(&Q, reader)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(Q.guess, 10.0, 1e-10) {
		t.Fatalf("Guess string not processed correctly: got %v, expected %s", Q.guess, guessStr)
	}

	// Test with invalid guess
	Q1 := NewQuestion(-50.0, 50.0)
	guessStr1 := "£)(*)£)(QSDLKJSLDJ\n"
	reader1 := bufio.NewReader(strings.NewReader(guessStr1))
	err1 := GetGuess(&Q1, reader1)
	if err1 == nil {
		t.Fatal("Expected an error for an invalid guess string.")
	}
}


func TestJudgeGuess(t *testing.T) {
	Q := NewQuestion(-10.0, 10.0)
	guess := Q.ans + 2.0
	guessStr := fmt.Sprintf("%f\n", guess)
	reader := bufio.NewReader(strings.NewReader(guessStr))
	GetGuess(&Q, reader)
	resMsg := JudgeGuess(Q)
	fmt.Print(resMsg)
	// TODO: Test that resMsg contains (some of?) the expected words.
}
