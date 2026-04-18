package main

import "fmt"

func main() {

	// 1. INITIALIZING A MAP (Map Literal)
	// map[KeyType]ValueType
	studentGrades := map[string]int{
		"Alice": 90,
		"James": 85,
		"Dan":   60,
	}
	fmt.Printf("%+v\n", studentGrades)

	// 2. UPDATING OR ADDING
	// If the key exists, it updates. If not, it creates a new entry.
	studentGrades["Alice"] = 95
	fmt.Printf("%+v\n", studentGrades)

	// 3. THE "COMMA OK" IDIOM (Very Important!)
	// In Go, if you look up a key that doesn't exist, it returns the "zero value" (0 for ints).
	// To tell the difference between a grade of 0 and a missing student, we use 'ok'.
	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Printf("Alice: %+v\n", alice) // ok is true if the key exists
	}

	// You can also do this in a single "if" block (short-statement syntax)
	key := "James"
	if value, ok := studentGrades[key]; ok {
		fmt.Printf("%s: %+v\n", key, value)
	}

	// 4. DELETING
	// Built-in delete function: delete(map, key)
	delete(studentGrades, "Alice")
	fmt.Printf("%+v\n", studentGrades)

	// 5. USING MAKE
	// Like slices, you can use make to initialize an empty map.
	configs := make(map[string]int)
	fmt.Printf("%+v %T\n", configs, configs)

	// 6. NIL VS EMPTY
	// A map created with 'make' is NOT nil. It is ready to use.
	// If you just declared 'var m map[string]int', it would be nil and 
	// trying to add a key to it would cause the program to crash (panic).
	if configs == nil {
		fmt.Printf("Config is nil")
	}
}
