package main

import "fmt"

func main() {
	// Init var a
	a := 30
	// Point b to the memory address of a
	b := &a
	// Assign value stored at address b
	c := *b

	// var b is same as a, but c is not.
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(c)
	a = 20
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(c)

	// Assigning value to a
	*b = 10
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(c)
}
