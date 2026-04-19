package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		// 1. THE PANIC
		// When panic is called, the current function stops immediately.
		// Go then starts "unwinding" the stack, running any 'defer' functions it finds.
		panic("something went wrong in mightPanic")
	}

	fmt.Println("This function executed without a panic")
}

func recoverable() {
	// 2. THE RECOVER (Inside a Defer)
	// 'recover' ONLY works when called inside a deferred function.
	// It's like catching a falling object before it hits the ground.
	defer func() {
		// r will be nil if there was no panic.
		// If there WAS a panic, r will contain the message from panic().
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// This call will trigger the panic logic.
	mightPanic(true)
	
	// This line will NEVER run because mightPanic stopped the flow here.
	fmt.Println("After mightPanic")
}

func main() {
	// Because 'recoverable' caught the panic, the program doesn't crash.
	// main() will finish gracefully.
	recoverable()
	fmt.Println("Main finished safely")
}
