package main

import (
	"fmt"
)

// 1. A BASIC FUNCTION
// Takes one "parameter" (name) which must be a string.
// It does its work and ends; it does NOT return a value.
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 2. MULTIPLE PARAMETERS
// When parameters are the same type, you could even write 'func add(a, b int)'.
// This function also prints its result directly rather than returning it.
func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

// 3. A FUNCTION WITH A RETURN VALUE
// Notice the 'float64' at the end of the line: '...height float64) float64 {'
// This tells Go: "I promise to give you back a decimal number (float64) when I'm done."
func calculateArea(width float64, height float64) float64 {
	// Guard Clause: It's common in Go to check for errors at the start.
	if width < 0 || height < 0 {
		fmt.Println("Error: width and height must be positive")
		// Since we promised to return a float64, we must return SOMETHING (0).
		return 0
	}
	
	// The 'return' keyword sends the result back to whoever called the function.
	return width * height
}

func main() {
	// Calling the greet function.
	greet("Bob Wonderland")
	
	// Calling the add function.
	add(1, 2)

	// Because calculateArea RETURNS a value, we can capture it in a variable (area).
	area := calculateArea(4, 4)
	fmt.Println(area)
}
