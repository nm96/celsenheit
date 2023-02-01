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


// RandFloat generates a random floating point number in the interval [min, max]
func RandFloat(min, max float64) float64 {
	return min + rand.Float64() * (max - min)
}


// ContainsStr returns true if a given string slice contains a given string.
func ContainsStr(list []string, element string) bool {
	for _, s := range list {
		if s == element {
			return true
		}
	}
	return false
}


// VerboseDegreeConversion verbosely outputs the results of a degree conversion
// (either F->C or C->F) to the command line.
func VerboseDegreeConversion(v float64, fromScale string, toScale string) (string, error) {
	// Define list of supported temperature scales
	scales := []string{"F","C"}

	// Check that input temperature scales are valid.
	if !ContainsStr(scales, fromScale) || !ContainsStr(scales, toScale) || fromScale == toScale {
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


// Define a new struct which holds all the information necessary to present and
// mark a question.
type Question struct {
	fromScale string
	toScale string
	val float64
	ans float64
	guess float64
}


// NewQuestion randomly generates a conversion question, computes the required
// result and returns this information in a Question struct.
func NewQuestion(Cmin, Cmax float64) Question {
	// Intialize variables for the scales and values.
	var fromScale, toScale string
	var val, ans float64

	// Randomly choose the conversion direction and the value to convert.
	switch rand.Intn(2) {
	case 0:
		fromScale, toScale = "C", "F"
		val = RandFloat(Cmin, Cmax)
		ans = C2F(val)

	case 1:
		fromScale, toScale = "F", "C"
		val = RandFloat(C2F(Cmin), C2F(Cmax))
		ans = F2C(val)
	}
	return Question{fromScale, toScale, val, ans, 0}
}


// GetGuess issues a question, gets a response from the command line and
// attempts to convert that to a guess value which is added to the Question
// struct, ready for marking.
func GetGuess(Q_p *Question, reader *bufio.Reader) error {
	// Attempt to read input string and convert it to a float.
	fmt.Printf("Convert %.3g\u00b0%s to \u00b0%s: ", Q_p.val, Q_p.fromScale, Q_p.toScale)

	// Attempt to read input string from command line.
	guessStr, readErr := reader.ReadString('\n')
	if readErr != nil {
		return readErr
	}

	// Quit session if user has typed "q", "Q", "quit" etc.
	if len(guessStr) > 0 && strings.ToUpper(guessStr[:1]) == "Q" {
		fmt.Println("Exiting.")
		fmt.Printf("(%.3g\u00b0%s is equivalent to %.3g\u00b0%s.)\n", Q_p.val,
		Q_p.fromScale, Q_p.ans, Q_p.toScale)
		os.Exit(0)
	}

	// Attempt to convert input string to float.
	guessStr = strings.TrimSpace(guessStr) // (Remove \n)
	guess, convErr := strconv.ParseFloat(guessStr, 64)
	if convErr != nil {
		return convErr
	}

	Q_p.guess = guess
	return nil
}


// JudgeGuess takes a question where the answer has been (presumable) guessed,
// and returns a string giving feedback about how close the guess was and what
// the correct answer is.
func JudgeGuess(Q Question) string {
	var guess, ans float64
	if Q.toScale == "F" {
		guess = F2C(Q.guess)
		ans = F2C(Q.ans)
	} else {
		guess = Q.guess
		ans = Q.ans
	}

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
	resMsg := ""
	gap := math.Abs(ans - guess)
	for i := 0; i < len(thresholds) + 1; i++ {
		if i == len(thresholds) || gap < thresholds[i] {
			resMsg += messages[i]
			resMsg += fmt.Sprintf(" You were off by %.4g\u00b0C: ", gap)
			break
		}
	}
	resMsg += fmt.Sprintf("%.3g\u00b0%s is equivalent to %.3g\u00b0%s.\n",
	Q.val, Q.fromScale, Q.ans, Q.toScale)
	return resMsg
}


// RunGuess combines the three functions above to create and issue a question
// and mark the result, which is printed to the console.
func RunGuess(reader *bufio.Reader) {
	// Define parameters and create new question.
	Cmin, Cmax := -50.0, 50.0
	Q := NewQuestion(Cmin, Cmax)
	// Get guess from command line, and keep repeating if guess is invalid.
	guessErr := GetGuess(&Q, reader)
	for guessErr != nil {
		guessErr = GetGuess(&Q, reader)
	}
	// Print feedback and correct answer.
	fmt.Print(JudgeGuess(Q))
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

		// Run RunGuess function repeatedly.
		for {
			reader := bufio.NewReader(os.Stdin)
			RunGuess(reader)
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
		resStr, err := VerboseDegreeConversion(v, fromScale, toScale)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(resStr)
		return
	} else {
		fmt.Println("Command-line arguments not understood.")
		fmt.Println(
			`Usage: celsenheit degree_value convert_from convert_to
			e.g. : celsenheit 20.0 C F`)
		return
	}
}
