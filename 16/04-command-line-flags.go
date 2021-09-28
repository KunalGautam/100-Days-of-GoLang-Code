package main

import (
	"flag"
	"fmt"
)

func main() {
	// Crete getWord flag, with default value Hello World and short description about flag
	flags := flag.String("getWord", "Hello World", "short description about flag")
	flag.Parse()
	fmt.Println("getWord Provided:", *flags)
}

// run go run 04-command-line-flags.go -h
// run go run 04-command-line-flags.go
// run go run 04-command-line-flags.go -getWord=Test
