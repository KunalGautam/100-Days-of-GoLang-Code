package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Create temp file
	file, err := os.CreateTemp("", "test.txt")
	if err != nil {
		panic(err)
	}

	// Write data to temp file
	_, err = file.Write([]byte("Hello"))
	if err != nil {
		panic(err)
	}

	// Read temp file
	rFile, err := os.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rFile))
	// get path of temp file
	pathOfFile, err := filepath.Abs(file.Name())
	fmt.Println(pathOfFile)

	// remove when program exists
	os.RemoveAll(file.Name())

}
