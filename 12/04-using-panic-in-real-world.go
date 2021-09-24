package main

import "os"

func init() {
	_, err := os.ReadFile("/file/not/present")
	if err != nil {
		panic(err)
	}
}

func main() {
	// Run code if file can be read.
}
