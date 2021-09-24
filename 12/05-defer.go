package main

import (
	"fmt"
)

func main() {
	defer one()
	two()

}

func one() {
	fmt.Println("Func One Ran")
}

func two() {
	fmt.Println("Func Two Ran")
}
