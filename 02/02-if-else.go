package main

import "fmt"

func main() {
	// Step 1
	if 5%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}

	// Step 2
	a := 8
	if a%3 == 0 {
		fmt.Println("Number is divisible by three")
	} else {
		fmt.Println("Number is Not divisible by three")
	}

}
