package main

import "fmt"

// 1. THE VARIADIC SUM FUNCTION
// The '...int' tells Go: "You can pass any number of integers here."
// Inside the function, 'numbers' is treated as a slice: []int
func sum(numbers ...int) int {
	total := 0
	// Since 'numbers' is a slice, we can use 'range' to loop through it.
	for _, number := range numbers {
		total += number
	}
	return total
}

// 2. OPTIONAL ARGUMENTS PATTERN
// Variadic functions are often used to handle "optional" settings.
func config(numbers ...int) {
	// We check the length to see if the user actually passed anything.
	if len(numbers) > 0 {
		// If they did, we take the first value.
		first := numbers[0]
		fmt.Println("First number:", first)
	} else {
		// If they passed nothing, we can provide a default behavior.
		fmt.Println("Default number")
	}
}

func main() {

	// You can pass as many arguments as you want.
	// Go "packs" these into a slice for the function.
	fmt.Println(sum(1, 2, 3, 4))

	// You can also pass NO arguments.
	// In this case, 'numbers' inside config() will be an empty slice (len 0).
	config() 

	// You can also pass a single argument.
	config(10)
}
