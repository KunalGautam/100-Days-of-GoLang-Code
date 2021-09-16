package main

import "fmt"

func main() {
	// Step 1: Sum
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//Step 2: Index Printing
	for i, num := range nums {
		if num == 4 {
			fmt.Println("index of 4 is at:", i)
		}
	}

}
