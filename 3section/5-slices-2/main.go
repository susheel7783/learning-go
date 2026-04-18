package main

import (
	"fmt"
	"slices" // This is a standard library package for slice utilities
)

func main() {

	fmt.Println("--- Advanced Slicing Operations ---")
	// Start with a slice of 10 numbers.
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v, len: %d, cap: %d\n", original, len(original), cap(original))
    // %v value %d int,%s string
	// 1. SLICING WITH [low:high]
	// Takes elements from index 2 up to (but NOT including) index 5.
	// Elements: original[2], original[3], original[4]
	s1 := original[2:5] 
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	// NOTE: len is 3, but cap is 8. 
	// Capacity starts from the START of the slice (index 2) to the end of the underlying array.

	// 2. OMITTING LOW [ :high]
	// If you leave out the first number, it defaults to 0.
	s2 := original[:4] 
	fmt.Printf("s2 (original[:4]): %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	// 3. OMITTING HIGH [low: ]
	// If you leave out the second number, it goes all the way to the end.
	s3 := original[6:]
	fmt.Printf("s3 (original[6:]): %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))

	// 4. THE FULL SLICE [:]
	// This creates a new slice header that points to the entire original array.
	s4 := original[:]
	fmt.Printf("s4 (original[:]): %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))

	// 5. UTILITY FUNCTIONS
	// slices.Contains checks if a value exists in the slice.
	// (Note: This line evaluates to true, but doesn't do anything because we aren't saving the result)
	slices.Contains(s4, 8)

	// 6. GROWING BEYOND CAPACITY
	// We append 1000 to s4. Since s4 was at its full capacity (10), 
	// Go must allocate a NEW underlying array with more room.
	s4 = append(s4, 1000)

	fmt.Printf("s4 (modified original[:]): %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))
	// Notice: len is now 11, and cap has doubled to 20!
}
