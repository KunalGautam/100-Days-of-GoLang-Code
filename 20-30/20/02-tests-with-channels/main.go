package main

import "fmt"

func main() {
	for i := range GenerateInts() {
		fmt.Printf("%v\n", i)
	}
}

func GenerateInts() <-chan int {
	rc := make(chan int)
	go func() {
		defer close(rc)

		for i := 0; i < 10; i++ {
			rc <- i
		}
	}()
	return rc
}
