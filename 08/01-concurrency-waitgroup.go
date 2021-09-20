package main

import (
	"fmt"
	"sync"
)

// Init wait group var
var wg sync.WaitGroup

func runThis(capture string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Ran from ", capture, ": ", i)
	}
	wg.Done()
}

func main() {
	//runThis("S1")

	// Provide number of go routines to wait for
	wg.Add(3)
	go runThis("S2")
	go runThis("S3")

	//Anonymous function
	go func(msg string) {
		fmt.Println("Hello from ", msg)
		wg.Done()
	}("S4")

	// As the above subroutines run separately, run this when all is done
	wg.Wait()
	fmt.Println("done")
}
