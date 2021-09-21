// In this example a single variable 'count' is accessed by multiple go routines.
// Although the output should be 4000, but due to race condition it will vary on multiple core processors.
// go run -race filename.go  | -race flag tests if there is race condition in the code or not.
package main

import "sync"

var count int
var wg sync.WaitGroup

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
		count++
	}
	wg.Done()
}
