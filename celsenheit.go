package main

import (
	"fmt"
	"strconv"
	"flag"
)


// F2C converts a temperature value from Fahrenheit to Celsius. 
func F2C(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) / 1.8
}


// C2F converts a temperature value from Celsius to Fahrenheit.
func C2F(celsius float64) float64 {
	return celsius * 1.8 + 32.0
}


// printDegreeConversion verbosely outputs the results of a degree conversion
// (either F->C or C->F) to the command line.
func printDegreeConversion(s string, toScale string) {
	var fromScale string
	var r float64

	switch toScale {
	case "C":
		fromScale = "F"
	case "F":
		fromScale = "C"
	default:
		fmt.Println("Scale (C or F) not provided for degree conversion")
		return
	}

	fmt.Printf("Converting %s\u00b0%s to \u00b0%s:\n", s, fromScale, toScale)

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Error attempting to convert %s to a real number.\n", s)
		fmt.Println(err)
		return
	} else {
		switch fromScale {
		case "C":
			r = C2F(v)
		case "F":
			r = F2C(v)
		}
	}

	fmt.Printf("%g\u00b0%s is equivalent to %.3g\u00b0%s.\n", v, fromScale, r, toScale)
}


func main() {
	// Declare command-line flags for converting C <-> F
	// (Flag is the symbol of the scale you want to convert TO)
	C2Fptr := flag.Bool("F", false, "Bool: convert from Celsius to Fahrenheit?")
	F2Cptr := flag.Bool("C", false, "Bool: convert from Fahrenheit to Celsius?")
	flag.Parse()

	// Convert from C to F if required.
	if *C2Fptr {
		Cstr := flag.Args()[0]
		printDegreeConversion(Cstr, "F")
	}

	// Convert from F to C if required
	if *F2Cptr {
		Fstr := flag.Args()[0]
		printDegreeConversion(Fstr, "C")
	}
}
