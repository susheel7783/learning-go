package main

import "fmt"

func main() {

	// 1. SLICE LITERAL
	// Unlike an array [3]string, this []string has no number inside the brackets.
	// This creates a slice that points to an underlying array of 3 names.
	names := []string{"Alice", "John", "Mark"}
	fmt.Println(names)

	// 2. THE MAKE FUNCTION
	// make([]type, length, capacity)
	// This creates a slice of 3 zeros, but pre-allocates space for 5.
	items := make([]int, 3, 5)

	// Len: The number of elements currently in the slice.
	// Cap: The total space available in the underlying array before it needs to grow.
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
	// Output: Items: [0 0 0], Len: 3, Cap: 5

	// 3. THE APPEND FUNCTION
	// append adds elements to the end. 
	items = append(items, 1) // Len 4
	items = append(items, 2) // Len 5 (Capacity is now full!)
	
	// WATCH THIS: When we add 3 and 4, the slice exceeds its capacity (5).
	// Go automatically creates a NEW, larger underlying array and copies the data over.
	items = append(items, 3) 
	items = append(items, 4) 
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
	// Output: Items: [0 0 0 1 2 3 4], Len: 7, Cap: 10 (Go usually doubles the capacity)

	// 4. SLICING A SLICE
	// items[3:7] means "start at index 3, up to (but NOT including) index 7".
	// This creates a 'view' of the same underlying data.
	fmt.Printf("%+v", items[3:7]) // Output: [1 2 3 4]

}
