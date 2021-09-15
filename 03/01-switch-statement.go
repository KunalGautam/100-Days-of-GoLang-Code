package main

import "fmt"

func main() {
	// Step1
	a := 3
	switch a % 2 {
	case 0:
		fmt.Println("Number is Even")
	case 1:
		fmt.Println("Number is Odd")
	}

	// Step 2: Set default
	switch a % 5 {
	case 0:
		fmt.Println("Number is divisible by 5")
	default:
		fmt.Println("Number is not divisible by 5")

	}

}
