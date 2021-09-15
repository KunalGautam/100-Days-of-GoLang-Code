package main

import "fmt"

func main() {
	// Step 1: Initialize slice with zero(Default) values
	// make() is a special built-in function that is used to initialize slices, maps, and channels.
	a := make([]int, 4)
	fmt.Println(a)
}
