package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"math/rand"
	"time"
	"log"
	"strings"
)


// Define usage help message.
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


// randFloat generates a random floating point number in the interval [max, min]
func randFloat(min, max float64) float64 {
	return min + rand.Float64() * (max - min)
}


// Contains returns true if a given string slice contains a given string.
func Contains(list []string, element string) bool {
	for _, s := range list {
		if s == element {
			return true
		}
	}
	return false
}


// verboseDegreeConversion verbosely outputs the results of a degree conversion
// (either F->C or C->F) to the command line.
func verboseDegreeConversion(v float64, fromScale string, toScale string) {
	// Define list of supported temperature scales
	scales := []string{"F","C"}

	// Check that input temperature scales are valid.
	if !Contains(scales, fromScale) || !Contains(scales, toScale) || fromScale == toScale {
		fmt.Printf("Invalid conversion %s->%s: Only F->C and C->F are currently supported.\n",
		fromScale, toScale)
		return
	}

	// Print the conversion about to be performed.
	fmt.Printf("Converting %g\u00b0%s to \u00b0%s:\n", v, fromScale, toScale)

	// Convert value using the specified conversion function.
	var result float64
	switch fromScale {
	case "C":
		result = C2F(v)
	case "F":
		result = F2C(v)
	}

	// TODO: Mark guess and issue feedback.

	// Print conversion result
	fmt.Printf("%g\u00b0%s is equivalent to %.3g\u00b0%s.\n", v, fromScale, result, toScale)
}


func runGuess() {
	// Define ranges for values to translate
	Cmin, Cmax := -50.0, 50.0
	Fmin, Fmax := -60.0, 120.0
	var fromScale, toScale string
	var val, ans float64
	switch rand.Intn(2) {
	case 0:
		fromScale, toScale = "C", "F"
		val = randFloat(Cmin, Cmax)
		ans = C2F(val)

	case 1:
		fromScale, toScale = "F", "C"
		val = randFloat(Fmin, Fmax)
		ans = F2C(val)
	}
	fmt.Printf("Convert %.3g\u00b0%s to \u00b0%s: ", val, fromScale, toScale)
	reader := bufio.NewReader(os.Stdin)
	guessStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	guessStr = strings.TrimSpace(guessStr) // Remove \n
	guess, err := strconv.ParseFloat(guessStr, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your guess:", guess)
	fmt.Printf("%.3g\u00b0%s is equivalent to %.3g\u00b0%s.\n", val, fromScale, ans, toScale)
}


func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		// Default behaviour: run the app in repeated guess mode.
		fmt.Println("Celsenheit guess mode: practice temperature conversions on random values!")
		fmt.Println("=========================================================================")
		fmt.Println()
		for {
			runGuess()
			fmt.Println()
		}
		return
	} else if len(os.Args) == 4 {
		// If given the right inputs, run app as a converter tool.

		// Take temperature scale strings from command line, normalizing to the
		// capitalized first letter of the string, e.g. celsius -> C.
		fromScale := strings.ToUpper(os.Args[2][:1])
		toScale := strings.ToUpper(os.Args[3][:1])

		// Take input temperature value string from command line and convert to
		// float if possible.
		valueString := os.Args[1]
		v, err := strconv.ParseFloat(valueString, 64)
		if err != nil {
			fmt.Printf("Input error: cannot convert %s to a temperature value.\n", valueString)
			return
		}

		// Run the verbose degree conversion with these sanitized inputs.
		verboseDegreeConversion(v, fromScale, toScale)
		return
	} else {
		fmt.Println("Command-line arguments not understood.")
		fmt.Println(usage)
		return
	}

}
