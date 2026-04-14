package main

import (
	"fmt"
	"strings"
)

// MathError is a custom "struct" (a collection of fields). 
// In Go, we create custom errors by making a struct and giving it an Error() method.
type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

// These are constants—unchanging values used to avoid "magic strings" in the code.
const (
	division       = "Division"
	divisionErrMsg = "division by zero is not allowed"
)

// This is the Error method. By defining this, MathError now "satisfies" 
// the built-in error interface, meaning we can use it anywhere an error is expected.
func (e MathError) Error() string {
	var inputs []string
	if e.Operation == division {
		// fmt.Sprintf creates a formatted string without printing it to the console.
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}

	return fmt.Sprintf("Math error in %s (%s): %s",
		e.Operation,
		strings.Join(inputs, ","),
		e.Message)
}

// Sum is a "variadic function" (the ... means it can take any number of integers).
func Sum(numbers ...int) int {
	// 'defer' tells Go to run this line at the very end of the function, 
	// just before it returns. It's great for cleanup tasks.
	defer fmt.Println("Sum finished")

	total := 0
	// 'range' lets us loop through the slice of numbers.
	// We use '_' because we don't need the index, only the value 'n'.
	for _, n := range numbers {
		total += n
	}
	return total
}

// SafeDivision returns two things: the result (int) and an error.
// Returning (value, error) is the standard way to handle problems in Go.
func SafeDivision(a, b int) (int, error) {
	if b == 0 {
		// We return a pointer (&) to our custom MathError struct.
		return 0, &MathError{
			Operation: division,
			InputA:    a,
			InputB:    b,
			Message:   divisionErrMsg,
		}
	}

	// If b is not 0, we return the result and 'nil' for the error.
	return a / b, nil
}

func main() {
	// Calling the variadic function with 3 arguments.
	fmt.Println(Sum(1, 2, 3))

	// In Go, we check if err != nil immediately after calling a function.
	value, err := SafeDivision(10, 0)
	if err != nil {
		// If there's an error, print it. This calls our custom Error() method automatically!
		fmt.Println(err)
	}

	fmt.Println("SafeDivision", value)
}
