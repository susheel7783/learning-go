package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	
	// This will not run now. It is "saved" for later.
	defer fmt.Println("Function simpleDefer: deferred")
	
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	// ... all the "Middle" prints happen first ...
	
	// RIGHT HERE, just before the function finishes, the deferred line runs.
}

func lifoSimpleDefer() {
	fmt.Println("Function lifoSimpleDefer: Start")
	
	// 1. LIFO (Last-In, First-Out)
	// When you have multiple defers, they are added to a "Stack".
	// The last one you write is the FIRST one to run at the end.
	defer fmt.Println("First: deferred")  // Added to stack first
	defer fmt.Println("Second: deferred") // Added to stack second
	
	fmt.Println("Function lifoSimpleDefer: Middle")
	
	// When this function ends:
	// "Second: deferred" prints first.
	// "First: deferred" prints second.
}

func main() {
	// 2. RESOURCE MANAGEMENT
	// We create a file. In other languages, if we had an error later, 
	// we might forget to close the file, causing a memory leak.
	file, err := os.Create("./defer.txt")
	if err != nil {
		fmt.Println(err)
	}
	
	// By putting 'defer' immediately after a successful file creation, 
	// we guarantee the file will close no matter how complex main() becomes.
	defer file.Close()

	lifoSimpleDefer()

	fmt.Println("Last in main()")
	
	// RIGHT HERE, file.Close() finally executes.
}
