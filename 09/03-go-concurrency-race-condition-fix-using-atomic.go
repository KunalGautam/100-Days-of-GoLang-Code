// In my understanding, the atomic package from sync can be used to fix race condition instead of mutex.
// atomic package is limited to int operations only.
package main

import (
	"sync"
	"sync/atomic"
)

var count int64
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
		// atomic
		atomic.AddInt64(&count, 1)
	}
	wg.Done()
}
