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


func main() {
	// Declare command-line flags for converting C <-> F
	// (Flag is the symbol of the scale you want to convert TO)
	C2Fptr := flag.Bool("F", false, "Bool: convert from Celsius to Fahrenheit?")
	F2Cptr := flag.Bool("C", false, "Bool: convert from Fahrenheit to Celsius?")
	flag.Parse()

	// Convert from C to F if required.
	if *C2Fptr {
		Cstr := flag.Args()[0]
		fmt.Printf("Converting %s degrees C to F\n", Cstr)
		C, err := strconv.ParseFloat(Cstr, 64)
		if err != nil {
			fmt.Printf("Error attempting to convert %s to a real number.\n", Cstr)
			fmt.Println(err)
		} else {
			F := C2F(C)
			fmt.Printf("%g degrees Celsius is equivalent to %g degrees Fahrenheit.\n", C, F)
		}
	}

	// Convert from F to C if required
	if *F2Cptr {
		Fstr := flag.Args()[0]
		fmt.Printf("Converting %s degrees F to C\n", Fstr)
		F, err := strconv.ParseFloat(Fstr, 64)
		if err != nil {
			fmt.Printf("Error attempting to convert %s to a real number.\n", Fstr)
			fmt.Println(err)
		} else {
			C := F2C(F)
			fmt.Printf("%g degrees Fahrenheit is equivalent to %g degrees Celsius.\n", F, C)
		}
	}
}
