package main

import (
	"fmt"
	"time"
)

func worker(done chan bool, s string) {
	fmt.Println("running ", s)
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 2)
	go worker(done, "s1")
	go worker(done, "s2")

	//Block until we receive a notification from the worker on the channel.
	<-done
}
