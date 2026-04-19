package main

import (
	"errors"
	"fmt"
	"time"
)

// 1. SENTINEL ERRORS
// These are "standard" error variables defined at the package level.
// We use these so we can compare errors later using errors.Is().
var ErrDivisionByZero = errors.New("division by zero")
var ErrNumTooLarge = errors.New("number too large")

// 2. CUSTOM ERROR TYPE
// Sometimes a simple string isn't enough. This struct stores the 
// operation name, an error code, a message, and the exact time it happened.
type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

// To make our struct an "error", it MUST have an Error() method.
func (op OpError) Error() string {
	return op.Message
}

// A "Constructor" function to make creating OpErrors easier.
func NewOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

func DoSomething() error {
	// We return our custom struct as an 'error' type.
	return NewOpError("doSomething", 100, "do something failed", time.Now())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero // Returning a sentinel
	}

	if a > 1000 {
		return 0, ErrNumTooLarge // Returning a sentinel
	}

	return a / b, nil
}

func main() {
	value, err := divide(1001, 1)

	if err != nil {
		// 3. THE ERRORS.IS() FUNCTION
		// This is the modern way to check which specific error occurred.
		// It works even if the error was "wrapped" inside another error.
		if errors.Is(err, ErrDivisionByZero) {
			fmt.Println("Handle: User tried to divide by zero")
		} else if errors.Is(err, ErrNumTooLarge) {
			fmt.Println("Handle: The input was way too big")
		}
		return
	}

	fmt.Println(value)
}
