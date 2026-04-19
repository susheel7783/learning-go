package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// 1. MULTIPLE RETURN VALUES (The idiomatic way)
// This function returns an (int) AND an (error).
// This is how Go handles problems instead of using "exceptions."
func divide(a, b int) (int, error) {
	if b == 0 {
		// errors.New creates a simple error object with a message.
		return 0, errors.New("divide by zero")
	}

	// If everything is fine, we return the result and 'nil' for the error.
	return a / b, nil
}

// 2. NAMED RETURN VALUES
// Here, we define the variable names (firstName, lastName) in the function signature.
// Go initializes these variables for us at the start of the function.
func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	
	// We assign values directly to the named return variables.
	firstName = parts[0]
	lastName = parts[1]

	// "Naked" Return: Because we named the returns in the signature, 
	// just typing 'return' automatically sends back firstName and lastName.
	return
}

func main() {

	// time.Now() is a built-in Go function that returns the current time.
	time.Now()

	// Handling Multiple Returns:
	// We must provide two variables (value, err) to "catch" the two things coming back.
	value, err := divide(10, 0)
	
	// The "if err != nil" pattern is the most common code block in Go.
	if err != nil {
		fmt.Println(err) // "divide by zero"
	} else {
		fmt.Println(value)
	}

	// Catching named return values works exactly the same way.
	firstName, lastName := splitName("Joseph Abah")

	fmt.Println(firstName, lastName)
}
