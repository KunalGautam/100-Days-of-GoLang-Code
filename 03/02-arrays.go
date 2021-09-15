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

}
