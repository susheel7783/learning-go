package main

import (
	"fmt"
)

// 1. THE BASIC INTERFACE
type Person interface {
	GetName() string
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

// 2. THE STRINGER INTERFACE
// Go has a built-in interface called 'Stringer' in the fmt package.
// It looks like this: type Stringer interface { String() string }
// By adding this method, you control what happens when you pass jane to fmt.Println.
func (e BusinessPerson) String() string {
	return fmt.Sprintf("Person[ID:%d, Name:%s]", e.ID, e.Name)
}

// 3. CUSTOM TYPES
// You are telling Go: "I want a new type called ID, but use an 'int' under the hood."
// This is great for making code more readable.
type ID int

// You can even attach methods to your custom type!
func (idx ID) String() string {
	return fmt.Sprintf("COMING FROM HERE ID: %d", idx)
}

func main() {

	jane := BusinessPerson{
		ID:   1,
		Name: "Jane",
	}

	// 4. PRINTING MAGIC
	// fmt.Println checks: "Does 'jane' have a String() method?"
	// Since it does, it uses your custom format instead of the default {1 Jane}.
	fmt.Println(jane)

	var myId ID
	myId = 30
	// Even though myId is basically an integer, because it has a String() method,
	// it prints your custom message!
	fmt.Println(myId)
}
