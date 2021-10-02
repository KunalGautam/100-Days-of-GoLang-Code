package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(0, 10, 29, 8))
}

func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
