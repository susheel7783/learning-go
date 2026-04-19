package main

import "fmt"

// 1. RECURSION
// A function that calls itself to solve a problem by breaking it into smaller pieces.
func factorial(n int) int {
	// Base Case: Without this, the function would call itself forever (and crash).
	if n <= 1 {
		return 1
	}
	// Recursive Step: It calls itself with (n-1).
	// For 5, it calculates: 5 * (4 * (3 * (2 * 1)))
	return n * factorial(n-1)
}

// 2. CLOSURES (Function Factories)
// This function returns ANOTHER function. 
// The returned function is a 'func() int'.
func intSeq() func() int {
	i := 0
	// This anonymous function "closes over" the variable 'i'.
	// Even after intSeq() finishes, this inner function keeps 'i' alive in memory.
	return func() int {
		i++
		return i
	}
}

func main() {
	// Testing Recursion: 5! = 5 * 4 * 3 * 2 * 1 = 120
	fmt.Println(factorial(5))

	// Testing Closures:
	// 'nextInt' now holds the inner function returned by intSeq.
	nextInt := intSeq()
	
	// Every time we call nextInt, it updates the SAME 'i' variable 
	// that it "captured" when it was created.
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	fmt.Println(nextInt()) // 4

	// 3. ANONYMOUS FUNCTIONS (Function Literals)
	// You can define a function right inside main and assign it to a variable.
	logger := func(msg string) {
		fmt.Println(msg)
	}

	logger("Hello World")
	logger("Hello World")
}
