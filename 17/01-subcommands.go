package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subCmd := flag.NewFlagSet("name", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'name' subcommand")
		os.Exit(1)
	}

	subCmd.Parse(os.Args[2:])
	fmt.Println("Sub Command ", subCmd.Args())
}

// go run <filename.go> name xyz
// Output:
// Sub Command  [xyz]
