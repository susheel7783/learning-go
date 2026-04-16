package main

import "fmt"

func main() {
	// 1. DECLARING AN ARRAY
	// This creates an array named 'numbers' that holds exactly 2 integers.
	// In Go, arrays are "zero-valued" by default, so this starts as [0, 0].
	var numbers [2]int
	fmt.Printf("%+v\n", numbers) // Output: [0 0]

	// 2. ASSIGNING VALUES
	// We use index numbers (starting at 0) to set values.
	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("%+v\n", numbers) // Output: [1 2]

	// 3. ARRAY LITERALS
	// You can declare and initialize an array at the same time using curly braces.
	primes := [4]int{2, 3, 5, 7}
	fmt.Printf("%+v\n", primes)
	
	// You can update a value at a specific index.
	primes[3] = 11
	fmt.Printf("%+v\n", primes) // Output: [2 3 5 11]

	// 4. ITERATING (LOOPING)
	// len(primes) returns the size (4). This loop prints each element on a new line.
	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d\n", primes[i])
	}

	// 5. MULTI-DIMENSIONAL ARRAYS
	// This creates a "matrix" or a grid. 
	// [2][3]int means 2 rows and 3 columns.
	var matrix [2][3]int
	matrix[0][0] = 1 // Top-left corner
	matrix[0][1] = 2 // First row, second column
	matrix[1][2] = 3 // Second row, third column (bottom-right)

	// Visualizing the matrix:
	// [1, 2, 0]
	// [0, 0, 3]
	fmt.Printf("%+v\n", matrix) 
}
