package main

import (
	"fmt"
	"os"
)

func main() {
	// Read File
	data, err := os.ReadFile("/etc/bashrc")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}

	// Write File
	writedata := []byte("Hello World!!!")
	err = os.WriteFile("/tmp/100daysofcode", writedata, 0644)
	if err != nil {
		panic(err)
	}

	readData, err := os.ReadFile("/tmp/100daysofcode")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(readData))
	}

	// Creating File
	createFile, err := os.Create("/tmp/dat2")
	if err != nil {
		panic(err)
	} else {
		createFile.Chmod(644)
		createFile.Write(writedata)
		createFile.Close()
	}

}
