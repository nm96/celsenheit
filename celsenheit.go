package main

import (
	"fmt"
	"os"
	"strconv"
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
	fmt.Printf("Converting %s degrees C to F\n", os.Args[1])
	C, err := strconv.ParseFloat(os.Args[1], 64) 
	if err != nil {
		fmt.Printf("Error attempting to convert %s to a real number.\n", os.Args[1])
		fmt.Println(err)
	} else {
		F := C2F(C)
		fmt.Printf("%g degrees Celsius is equivalent to %g degrees Fahrenheit.\n", C, F)
	}
}
