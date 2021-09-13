package main

import "fmt"

func main() {
	// Assigning type int
	var a, b int = 1, 2

	//Assigning int (Auto detected by go with value provided)
	c := 3

	// Assigning auto-detection of type by go
	d, e, f := true, false, "Hola!"

	fmt.Println(a, b, c, d, e, f)
}
