package main

import (
	"fmt"
)


// Celsius converts a temperature value from Fahrenheit to Celsius. 
func Celsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) / 1.8
}


// Fahrenheit converts a temperature value from Celsius to Fahrenheit.
func Fahrenheit(celsius float64) float64 {
	return celsius * 1.8 + 32.0
}

func main() {
	C := 20.0
	F := Fahrenheit(C)
	fmt.Printf("%v degrees Celsius is equivalent to %v degrees Fahrenheit.\n", C, F)
}
