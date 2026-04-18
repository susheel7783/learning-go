package main

import "fmt"

// 1. PASS BY VALUE
// When you call this, Go makes a complete COPY of the number.
// The 'val' inside this function is a different piece of memory than the 'num' in main.
func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue: %+v\n", val)
}

// 2. PASS BY REFERENCE (Using Pointers)
// The '*' in '*int' means this function expects a memory address, not a plain number.
func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	// The '*' in '*val' is "dereferencing." 
	// It means: "Go to the address stored in 'val' and change the actual number there."
	*val = *val * 10 
	fmt.Printf("modifyPointer: %+v\n", val) // This prints the memory address (e.g., 0xc0000120b0)
}

func main() {
	num := 10

	// This won't change 'num' because we only gave the function a copy.
	modifyValue(num)
	fmt.Println(num) // Output: 10

	// The '&' symbol is the "Address Of" operator. 
	// We are sending the actual location of 'num' to the function.
	modifyPointer(&num)
	fmt.Println(num) // Output: 100 (It changed!)

	// 3. POINTER VARIABLES
	grade := 50
	// gradePtr stores the memory address of grade.
	gradePtr := &grade
	
	fmt.Printf("gradePtr address: %+v\n", gradePtr) // Prints the location
	
	// This looks complex, but it's just showing that you can point to pointers.
	// *(&gradePtr) is essentially saying: "Get the address of the pointer, then look inside it."
	fmt.Printf("gradePtr value: %+v\n", *(&gradePtr)) 
}
