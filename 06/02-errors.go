package main

import (
	"errors"
	"fmt"
)

func main() {
	// create an error
	err := errors.New("custom error")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Type %T", err)
	}
}
