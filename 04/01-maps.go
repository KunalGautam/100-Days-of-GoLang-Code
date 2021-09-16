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

	// Step 3: Retreive a key
	fmt.Println(a["age"])

	// Step 4: Delete a key
	delete(a, "age")
	fmt.Println(a)

	// Step 5: Check if key present.
	keyVal, ifPresent := a["notpresent"]
	fmt.Println(keyVal)
	fmt.Println(ifPresent)

	keyVal1, ifPresent1 := a["height"]
	fmt.Println(keyVal1)
	fmt.Println(ifPresent1)

}
