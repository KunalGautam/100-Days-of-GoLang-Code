package main

import "fmt"

func main() {
	// Step1: Create channel using make()
	ch1 := make(chan string)
	go func() { ch1 <- "ping" }()
	go func() { ch1 <- "pong" }()

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

}
