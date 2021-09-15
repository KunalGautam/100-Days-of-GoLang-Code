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
	// Step 3: Append slice, which is not possible with arrays
	a = append(a, 10)
	fmt.Println(a)
	// Step 4: Copy slice
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(b)

	// Step 5: Seek slice data
	// Start from index 0 till 2
	fmt.Println(a[:2])
	//start from index 2 till 4
	fmt.Println(a[2:4])
	//start from index 2 till end
	fmt.Println(a[2:])

	//Step 6: short hand declaration, empty slice
	c := []string{"a", "b", "c"}
	d := []string{}
	fmt.Println(c)
	fmt.Println(d)

	// Step 7: Two Dimentional Slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Step 8: retreive values by iteration
	for i, v := range c { //range returns both the index and value
		fmt.Printf("%d element of c is %s\n", i, v)
	}
}
