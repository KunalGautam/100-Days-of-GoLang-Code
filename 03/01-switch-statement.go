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

}
