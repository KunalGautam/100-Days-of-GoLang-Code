package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println("Timer 1 fired at", <-timer1.C)

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped ", stop)
	}

	timer3 := time.NewTimer(time.Second * 2)
	go func() {
		fmt.Println("Timer 3 fired at", <-timer3.C)
	}()

	//Wait for timer3
	time.Sleep(3 * time.Second)
}
