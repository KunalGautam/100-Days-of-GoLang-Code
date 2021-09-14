package main

import "fmt"

func main() {
	// Method 1
	i := 0
	// While True, run Forever
	for {
		if i <= 10 {
			fmt.Println(i)
		} else {
			break
		}
		i++
	}

	// Method 2
	for a := 0; a <= 10; a++ {
		fmt.Println("a-> ", a)
	}
}
