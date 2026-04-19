package main

import (
	"fmt"
)
// an interface in Go is a way to define a contract of behavior. It tells the world: "I don't care what you are, as long as you can do these specific things."

// While a struct defines what an object is (its data), an interface defines what an object does (its methods).
// 1. THE INTERFACE DEFINITION
// An interface is a "Contract." It does not store data (like fields).
// Instead, it lists "Method Signatures." 
// Any type that provides these methods is said to "implement" or "satisfy" the interface.
type Person interface {
	// This is a requirement: "To be a Person, you must have a GetName() method 
	// that takes no arguments and returns a string."
	GetName() string
}

// 2. CONCRETE TYPES
// These are standard structs that hold data.
type Employee struct {
	ID   int
	Name string
}

type BusinessPerson struct {
	ID   int
	Name string
}

// 3. IMPLICIT IMPLEMENTATION
// Go does not use an "implements" keyword.
// By simply defining this method, BusinessPerson now automatically 
// satisfies the Person interface.
func (e BusinessPerson) GetName() string {
	return e.Name
}

// Employee also satisfies the Person interface.
func (e Employee) GetName() string {
	return e.Name
}

// 4. POLYMORPHISM IN ACTION
// This function doesn't care about the underlying type (Employee vs BusinessPerson).
// It only cares that the input 'p' follows the "Person" contract.
func displayPerson(p Person) {
	// Because p is a Person, Go guarantees it has a GetName() method.
	fmt.Println(p.GetName())
}

func main() {
	joe := Employee{
		ID:   1,
		Name: "Joe",
	}

	jane := BusinessPerson{
		ID:   1,
		Name: "Jane",
	}

	// We can pass different types into the same function!
	displayPerson(joe)
	displayPerson(jane)

	fmt.Println(jane)
}
