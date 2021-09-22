// The example code at https://gobyexample.com/tickers is best to understand in using channels along with tickers
package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
