// Defer will run if specified before error/panic
package main

import (
	"fmt"
)

func main() {
	defer one()
	panic("I'm panicing")
	two()

}

func one() {
	fmt.Println("Func One Ran")
}

func two() {
	fmt.Println("Func Two Ran")
}
