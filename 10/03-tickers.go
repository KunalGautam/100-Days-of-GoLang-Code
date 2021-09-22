package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// To run a job every second.
	ticker := time.NewTicker(time.Second)

	for i := 1; i <= 5; i++ {
		<-ticker.C
		fmt.Println("Running for " + strconv.Itoa(i) + "th time")
	}

}
