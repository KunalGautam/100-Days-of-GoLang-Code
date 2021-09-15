package main

import (
	"fmt"
)

func main() {
	//Step 1: Declaring Arrays and assign zero values (default values)
	var a [3]int
	fmt.Println(a)

	// Step 2: Assigning values
	a[0] = 101
	fmt.Println(a)
	fmt.Println(a[0])

	// Step 3: Shorthand declaration
	b := [4]int{1011, 18641, 5432, 353} // short hand declaration to create array
	fmt.Println(b)
	fmt.Println(b[3])
}
