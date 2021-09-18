package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// create an error
	err := errors.New("custom error")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Type %T", err)
	}

	// Step 2, real world error example
	_, err = os.Open("filename.txt")
	if err != nil {
		fmt.Println(err)
	}
}
