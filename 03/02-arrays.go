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

	// Step 4: Declare on fly
	c := [...]int{1, 4, 2, 5, 7, 8, 3}
	fmt.Println(c)

	// Step 5: Length of array
	fmt.Println(len(c))
	// Step 6: print last element
	fmt.Println(c[len(c)-1])
	//Step 7: Iteration
	for i, v := range c { //range returns both the index and value
		fmt.Printf("%d element of c is %d\n", i, v)
	}

}
