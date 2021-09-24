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

}
