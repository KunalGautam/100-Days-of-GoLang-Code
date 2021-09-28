package main

import "os"

func main() {
	// Create directory, subdir
	err := os.Mkdir("test", 0755)
	if err != nil {
		panic(err)
	}
	// Rename
	err = os.Rename("test", "test1")
	if err != nil {
		panic(err)
	}
	// Delete Directory
	err = os.Remove("test1")
	if err != nil {
		panic(err)
	}
	// Crete sub directories
	err = os.MkdirAll("test2/hello/world", 0755)
	if err != nil {
		panic(err)
	}
	// Delete Directory and subdir
	err = os.RemoveAll("test2")
	if err != nil {
		panic(err)
	}
}
