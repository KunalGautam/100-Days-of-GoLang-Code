package main

import "fmt"

func main() {
	//Step 1: Sum function
	fmt.Println(sum(1, 2))

	//Step 2: Sum all
	fmt.Println(sumAll(1, 2, 3, 4))

	// Step 3:assigning value to variable
	a := sumAll(2, 9, 56, 11)
	fmt.Println(a)

	// Step 4: Return multiple value
	fmt.Println(textAndSum("The sum is: ", 1, 2, 3, 4))
}

//sum accepts two int and return a int type
func sum(x int, y int) int {
	return x + y
}

func sumAll(x ...int) int {
	sum := 0
	for _, num := range x {
		sum += num
	}
	return sum
}

func textAndSum(s string, i ...int) (string, int) {
	sum := 0
	for _, num := range i {
		sum += num
	}
	return s, sum
}
