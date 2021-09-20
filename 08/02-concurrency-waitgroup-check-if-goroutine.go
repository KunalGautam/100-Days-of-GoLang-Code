package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Init wait group var
var wgg sync.WaitGroup

func runAndCheck(capture string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Ran from ", capture, ": ", i)
	}

	// Check if run from go routine, https://stackoverflow.com/a/56702614
	// If check is not done, the function can throw error at wgg.done() point if not run by go routine.
	count := runtime.Callers(3, make([]uintptr, 1))
	if count == 0 {
		wgg.Done()
	}
}

func main() {
	runAndCheck("S1")

	// Provide number of go routines to wait for
	wgg.Add(3)
	go runAndCheck("S2")
	go runAndCheck("S3")

	//Anonymous function
	go func(msg string) {
		fmt.Println("Hello from ", msg)
		wgg.Done()
	}("S4")

	// As the above subroutines run separately, run this when all is done
	wgg.Wait()
	fmt.Println("done")
}
