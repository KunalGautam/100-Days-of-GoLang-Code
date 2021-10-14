package main

import (
	"flag"
	"fmt"
	"github.com/gosuri/uilive"
	"os"
	"time"
)

func main() {
	flagTimer := flag.Int("time", 20, "Set timer in minutes.")
	flag.Parse()

	if *flagTimer < 1 {
		fmt.Println("Time should be greater than 0")
		os.Exit(0)
	}

	var minString string
	if *flagTimer == 1 {
		minString = "Minute"
	} else {
		minString = "Minutes"
	}
	fmt.Printf("Time Set for Pomodoro at %d %s\r\n", *flagTimer, minString)
	secondsElapsed := *flagTimer * 60

	toggle := "To Work"

	renderLiveScreen := uilive.New()
	renderLiveScreen.Start()

	for {
		if secondsElapsed < 1 {
			secondsElapsed = *flagTimer * 60
			if toggle == "To Work" {
				toggle = "For Break"
			} else {
				toggle = "To Work"
			}
		}
		_, _ = fmt.Fprintf(renderLiveScreen, "Remaining Time %s %s\n", toggle, time.Second*time.Duration(secondsElapsed))
		secondsElapsed -= 1
		time.Sleep(time.Second)

	}

	renderLiveScreen.Stop()

}
