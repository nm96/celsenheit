package main

import (
	"fmt"
	"strconv"
	"os"
)

var usage string =
`Usage: celsenheit degree_value convert_from convert_to
e.g. : celsenheit 20.0 C F`


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
func printDegreeConversion(s string, fromScale string, toScale string) {

	// Check that temperature scales are valid.
	if (toScale != "C") && (toScale != "F") {
		fmt.Println(toScale, "is not a valid temperature scale to convert to.")
		return
	}
	if (fromScale != "C") && (fromScale != "F") {
		fmt.Println(fromScale, "is not a valid temperature scale to convert from.")
		return
	}

	// Print the conversion about to be performed.
	fmt.Printf("Converting %s\u00b0%s to \u00b0%s:\n", s, fromScale, toScale)

	// Convert input value from string to float if possible, or print error.
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(s, "is not a valid temperature value.")
		fmt.Println(err)
		return
	}

	// Convert value and print result.
	var r float64 // Conversion result variable
	switch fromScale {
	case "C":
		r = C2F(v)
	case "F":
		r = F2C(v)
	}
	fmt.Printf("%g\u00b0%s is equivalent to %.3g\u00b0%s.\n", v, fromScale, r, toScale)
}


func main() {
	if len(os.Args) < 4 {
		fmt.Println("Not enough command line arguments.")
		fmt.Println(usage)
		return
	} else if len(os.Args) > 4 {
		fmt.Println("Too many command line arguments.")
		fmt.Println(usage)
		return
	} else if os.Args[2] == os.Args[3] {
		fmt.Println("Temperature scales must be different.")
		fmt.Println(usage)
		return
	}
	valueString := os.Args[1]
	fromScale := os.Args[2]
	toScale := os.Args[3]
	printDegreeConversion(valueString, fromScale, toScale)
}
