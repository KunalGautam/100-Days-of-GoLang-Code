package main

import "sync"

var count int
var wg sync.WaitGroup

// Create Mutex
var mu sync.Mutex

func main() {
	wg.Add(4)
	go runMe()
	go runMe()
	go runMe()
	go runMe()
	wg.Wait()
	println(count)

}

func runMe() {
	for i := 0; i < 1000; i++ {
		// Lock and unlock variable.
		mu.Lock()
		count++
		mu.Unlock()
	}
	wg.Done()
}
