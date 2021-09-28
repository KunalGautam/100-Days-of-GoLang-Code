package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// Take gocommand along with arguments too
	fmt.Println(args)
	// Take only arguments passed
	fmt.Println(args[1:])
}
