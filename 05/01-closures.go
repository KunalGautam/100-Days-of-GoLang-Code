// This example shows closures.
package main

import (
	"fmt"
)

//Create a function nextInt and whenever it will be assigned to a variable, it will be a func().
func nextInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// Init the var, it will convert variable to a func
	a := nextInt()
	//Print
	fmt.Println(a())
	// Value of i will persist
	fmt.Println(a())
	fmt.Println(a())

	// Start new var
	b := nextInt()
	fmt.Println(b())

	// Even though new var is init, it will remember the value of i.
	fmt.Println(a())

}
