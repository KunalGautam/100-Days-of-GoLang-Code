package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("!")

	//Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.
	//	For portability, the status code should be in the range [0, 125].
	os.Exit(3)

	fmt.Println("Never gonna show this!")
}
