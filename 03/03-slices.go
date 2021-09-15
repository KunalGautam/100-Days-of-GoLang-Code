package main

import "fmt"

func main() {
	// Step 1: Initialize slice with zero(Default) values
	// make() is a special built-in function that is used to initialize slices, maps, and channels.
	a := make([]int, 4)
	fmt.Println(a)
	// Step 2: Set values to slice
	a[0] = 2
	a[1] = 4
	a[2] = 6
	a[3] = 8
	fmt.Println(a)
}
