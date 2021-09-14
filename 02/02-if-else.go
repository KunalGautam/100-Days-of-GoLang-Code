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

	//Step 3 : Multiple if else
	if num := 11; num < 0 {
		fmt.Println(num, " is Negative")
	} else if num >= 10 {
		fmt.Println(num, " is more than or equals to 10")
	} else {
		fmt.Println(num, " is less than 10")
	}

}
