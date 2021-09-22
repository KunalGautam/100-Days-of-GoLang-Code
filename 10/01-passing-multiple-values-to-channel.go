package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type logger struct {
	seq     int
	message string
	time    time.Time
}

func main() {
	logData := make(chan logger)

	go func() {
		for i := 1; i < 100; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			logData <- logger{
				seq:     i,
				message: "Message at " + strconv.Itoa(i),
				time:    time.Now(),
			}
		}
		close(logData)

	}()

	for r := range logData {
		fmt.Println(r)
	}

}
