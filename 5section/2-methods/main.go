package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

// 1. POINTER RECEIVER METHODS
// The (e *Employee) part is the "Receiver". 
// It attaches this function to the Employee struct.
// Because it uses a pointer (*), it can actually modify the original struct.

func (e *Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) Deactivate() {
	// This changes the IsActive field on the original jane/joe.
	e.IsActive = false
}

func (e *Employee) SetJoinedAt(t time.Time) {
	e.JoinedAt = t
}

// 2. CONSTRUCTOR (Review)
func NewEmployee(id int, firstName, lastName, position string, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

// 3. REGULAR FUNCTION (Comparison)
// This does the same thing as the Deactivate() method, 
// but you have to pass the employee as an argument: deactivate(&jane).
func deactivate(e *Employee) {
	e.IsActive = false
}

func main() {
	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    1000,
		IsActive:  true,
	}

	fmt.Printf("%+v\n", jane)
	
	// Calling the method! Much cleaner than deactivate(&jane).
	jane.Deactivate() 
	fmt.Printf("%+v\n", jane)

	// Demonstrating time manipulation
	jane.SetJoinedAt(time.Now().Add(100000000 * time.Minute))
	fmt.Printf("%+v\n", jane)

	// 4. THE "NIL POINTER" TRAP
	// This is a common beginner mistake.
	joe := &Employee{} // Points to an empty employee
	joe = nil          // Now joe points to NOTHING (nil)

	// DANGER: Calling a method on a nil pointer will cause a "Panic" (crash)
	// because Go tries to look up 'FirstName' at a memory address that doesn't exist.
	// joe.FullName() // Uncommenting this would crash the program.
}
