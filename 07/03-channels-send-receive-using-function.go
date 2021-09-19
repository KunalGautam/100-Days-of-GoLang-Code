package main

import "fmt"

func myfunc(ch chan string) {
	fmt.Println(<-ch)
}
func main() {
	fmt.Println("Start")
	// Creating a channel
	ch := make(chan string)

	go myfunc(ch)
	ch <- "Run this"
	fmt.Println("End")
}
