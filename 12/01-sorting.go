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

	// Step3: check if sorted
	output := sort.IntsAreSorted(ints)
	fmt.Println(output)
	ints2 := []int{8, 98, 4, 176, 5}
	output2 := sort.IntsAreSorted(ints2)
	fmt.Println(output2)
	output3 := sort.StringsAreSorted(str)
	fmt.Println(output3)
	str2 := []string{"z", "c", "w", "a"}
	output4 := sort.StringsAreSorted(str2)
	fmt.Println(output4)

}
