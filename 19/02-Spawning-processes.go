package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	// Binary location
	fmt.Println(dateCmd)
	// Execute command
	output, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))

	// Create ls command
	ls := exec.Command("ls", "-la")
	// Create a grep command and search for .go files
	grep := exec.Command("grep", "\\.go")

	// Set grep's stdin to the output of the ls command.
	grep.Stdin, err = ls.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}

	// Set grep's stdout to os.Stdout
	grep.Stdout = os.Stdout

	// Start the grep command first. (The order will be last command first)
	err = grep.Start()
	if err != nil {
		log.Fatalln(err)
	}

	// Run the ls command. (Run calls start and also calls wait)
	err = ls.Run()
	if err != nil {
		log.Fatalln(err)
	}

	// Wait for the grep command to finish.
	err = grep.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}
