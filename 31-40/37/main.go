// Panic and resume...
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting.......")
	go panicRun()
	i := 0
	for {
		i++
		time.Sleep(time.Second)
		fmt.Println("Ran for ", i, " Seconds")
	}
}

func panicRun() {
	defer panicRun()
	if r := recover(); r != nil {
		fmt.Println("Recovering from: ", r)
	}
	fmt.Println("Running Panic in 5 Seconds")
	time.Sleep(time.Second * 5)
	panic("Paincking..........")
}
