package main

import (
	"fmt"
	"time"
)

func runThis(capture string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Ran from ", capture, ": ", i)
	}
}

func main() {
	runThis("S1")

	go runThis("S2")
	go runThis("S3")

	//Anonymous function
	go func(msg string) {
		fmt.Println("Hello from ", msg)
	}("S4")

	// As the above subroutines run separately, we need to wait them for finish
	time.Sleep(time.Second)
	fmt.Println("done")
}
