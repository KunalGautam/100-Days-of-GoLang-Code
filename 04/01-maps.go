package main

import "fmt"

func main() {
	// Step 1: Init a map (A map is key pair value)
	a := make(map[string]int)
	a["age"] = 20
	a["height"] = 170
	fmt.Println(a)

	// Step 2: length of map
	fmt.Println(len(a))

}
