package main

import (
	"fmt"
	"strconv"
	"os"
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
func printDegreeConversion(s string, fromScale string, toScale string) {
	fmt.Printf("Converting %s\u00b0%s to \u00b0%s:\n", s, fromScale, toScale)
	var r float64 // Conversion result
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
	valueString := os.Args[1]
	fromScale := os.Args[2]
	toScale := os.Args[3]
	printDegreeConversion(valueString, fromScale, toScale)
}
