package main

import (
	"fmt"
	"errors"
	"strconv"
	"strings"
	"os"
	"time"
	"bufio"
	"math"
	"math/rand"
)


// F2C converts a temperature value from Fahrenheit to Celsius. 
func F2C(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) / 1.8
}


// C2F converts a temperature value from Celsius to Fahrenheit.
func C2F(celsius float64) float64 {
	return celsius * 1.8 + 32.0
}


// randFloat generates a random floating point number in the interval [min, max]
func randFloat(min, max float64) float64 {
	return min + rand.Float64() * (max - min)
}


// containsStr returns true if a given string slice contains a given string.
func containsStr(list []string, element string) bool {
	for _, s := range list {
		if s == element {
			return true
		}
	}
	return false
}


// verboseDegreeConversion verbosely outputs the results of a degree conversion
// (either F->C or C->F) to the command line.
func verboseDegreeConversion(v float64, fromScale string, toScale string) (string, error) {
	// Define list of supported temperature scales
	scales := []string{"F","C"}

	// Check that input temperature scales are valid.
	if !containsStr(scales, fromScale) || !containsStr(scales, toScale) || fromScale == toScale {
		errMsg := fmt.Sprintf("Invalid conversion %s->%s: Only F->C and C->F are currently supported.\n", fromScale, toScale)
		return "", errors.New(errMsg)
	}

	// Convert value using the specified conversion function.
	var result float64
	switch fromScale {
	case "C":
		result = C2F(v)
	case "F":
		result = F2C(v)
	}

	// Print conversion result
	resultMessage := fmt.Sprintf("%g\u00b0%s is equivalent to %.3g\u00b0%s.\n",
	v, fromScale, result, toScale)
	return resultMessage, nil
}


func runGuess() {
	// Define ranges for values to translate
	Cmin, Cmax := -50.0, 50.0
	Fmin, Fmax := -60.0, 120.0

	// Intialize variables for the scales and values.
	var fromScale, toScale string
	var val, ans float64

	// Randomly choose the conversion direction and the value to convert.
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

	// Attempt to read input string and convert it to a float.
	fmt.Printf("Convert %.3g\u00b0%s to \u00b0%s: ", val, fromScale, toScale)
	reader := bufio.NewReader(os.Stdin)

	// Attempt to read input string from command line.
	guessStr, readErr := reader.ReadString('\n')

	// Quit session if user has typed "q", "Q", "quit" etc.
	if len(guessStr) > 0 && strings.ToUpper(guessStr[:1]) == "Q" {
		fmt.Println("Exiting.")
		os.Exit(0)
	}

	// Attempt to convert input string to float.
	guessStr = strings.TrimSpace(guessStr) // (Remove \n)
	guess, convErr := strconv.ParseFloat(guessStr, 64)

	// Repeat the prompt and start again if there was an error reading or
	// processing in the lines above. 
	for readErr != nil || convErr != nil {
		fmt.Printf("Convert %.3g\u00b0%s to \u00b0%s: ", val, fromScale, toScale)
		reader := bufio.NewReader(os.Stdin)

		// Attempt to read input string from command line.
		guessStr, readErr = reader.ReadString('\n')

		// Quit session if user has typed "q", "Q", "quit" etc.
		if len(guessStr) > 0 && strings.ToUpper(guessStr[:1]) == "Q" {
			fmt.Println("Exiting.")
			os.Exit(0)
		}

		// Attempt to convert input string to float.
		guessStr = strings.TrimSpace(guessStr) // (Remove \n)
		guess, convErr = strconv.ParseFloat(guessStr, 64)
	}

	// Mark guess and issue feedback.
	judgeGuess(guess, ans, toScale)

	// Print correct result of conversion.
	fmt.Printf("%.3g\u00b0%s is equivalent to %.3g\u00b0%s.\n", val, fromScale, ans, toScale)
}


func judgeGuess(guess, ans float64, toScale string) {
	// Convert guess and ans values to degrees C if necessary for consistency.
	if toScale == "F" {
		guess = F2C(guess)
		ans = F2C(ans)
	}

	// Initialise slice with thresholds and associated feedback messages for
	// judging guesses. len(messages) should always equal len(threholds) + 1.
	thresholds := []float64{0.3,
						   1.0,
						   3.0,
						   10.0}
	messages := []string{"Astonishing!",
						 "Very close!",
						 "Pretty close!",
						 "In the ballpark!",
						 "Better luck next time!"}

	// Iterate through thresholds and give feedback based on how close the guess
	// is to the correct answer.
	gap := math.Abs(ans - guess)
	for i := 0; i < len(thresholds) + 1; i++ {
		if i == len(thresholds) || gap < thresholds[i] {
			fmt.Printf("%s You were off by %.4g\u00b0C: ", messages[i], gap)
			break
		}
	}
}


func main() {
	if len(os.Args) == 1 {
		// Default behaviour: run the app in repeated guess mode.

		// Initialise random seed and print intro banner.
		rand.Seed(time.Now().UnixNano())
		introMsg :=
`
Celsenheit guess mode: practice mental conversion of temperature values
=======================================================================

Enter 'Q' to exit.

`
		fmt.Print(introMsg)

		// Run runGuess function repeatedly.
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
		fmt.Print(verboseDegreeConversion(v, fromScale, toScale))
		return
	} else {
		fmt.Println("Command-line arguments not understood.")
		fmt.Println(
			`Usage: celsenheit degree_value convert_from convert_to
			e.g. : celsenheit 20.0 C F`)
		return
	}

}
