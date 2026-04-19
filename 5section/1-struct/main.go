package main

import (
	"fmt"
	"time"
)

// 1. THE STRUCT DEFINITION
// Think of this as a schema or a template. 
// Every 'Employee' will have these specific "fields."
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

// 2. A CONSTRUCTOR FUNCTION
// Go doesn't have "constructors" built-in. Instead, we write a regular 
// function that takes arguments and returns a new instance of the struct.
func NewEmployee(id int, firstName, lastName, position string, isActive bool) Employee {
	// We use the field names to set the values.
	// JoinedAt is set automatically using time.Now().
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

func main() {

	// 3. INITIALIZING MANUALLY
	// This is a "struct literal." We are filling out the form ourselves.
	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    1000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	// 4. ACCESSING FIELDS
	// We use the "dot notation" (.) to read specific data from the struct.
	fmt.Println(jane.ID)
	fmt.Println(jane.FirstName)

	// 5. USING THE CONSTRUCTOR
	joe := NewEmployee(1, "John", "Doe", "Jane", true)

	fmt.Println(joe.FirstName)

	// 6. MODIFYING FIELDS
	// You can change a value just like a normal variable.
	joe.Salary = 100000000
	fmt.Println(joe.Salary)

	// 7. STRUCT POINTERS
	// joePtr points to the memory address where 'joe' lives.
	joePtr := &joe
	
	// IMPORTANT GO FEATURE: 
	// In C++, you would have to use -> to access fields via a pointer.
	// In Go, you can still use the dot (.)—Go automatically "dereferences" it for you!
	joePtr.IsActive = true
	joePtr.LastName = "John Adam"
	
	fmt.Println(joe) // You'll see the LastName change in the original 'joe'
}
