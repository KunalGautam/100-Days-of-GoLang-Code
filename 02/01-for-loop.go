package main

import "fmt"

func main() {
	i := 0
	// While True, run Forever
	for {
		if i <= 10 {
			fmt.Println(i)
		} else {
			break
		}
		i++
	}
}
