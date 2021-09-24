package main

import (
	"fmt"
	"sort"
)

func main() {
	//Step 1: Create slice of strings
	str := []string{"z", "c", "w", "a"}
	// Sort strings
	sort.Strings(str)
	fmt.Println(str)

	// Step 2: create slice of ints
	ints := []int{8, 98, 4, 176, 5}
	// Sort ints
	sort.Ints(ints)
	fmt.Println(ints)

}
