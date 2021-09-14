package main

import (
	"fmt"
	"time"
)

func main() {
	// Step 1 : Now returns the current local time.
	fmt.Println(time.Now())
	// Step 2 : Get Current Month, Year, Week, Day, date, hour, minutes, seconds of current Local Time
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())

	//Step 3: Timezone
	loc, err := time.LoadLocation("Asia/Thimphu")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(time.Now().In(loc))
	}

	// Step 4: Sleep in seconds
	time.Sleep(2 * time.Second)
	fmt.Println("I'm awake now!")

	// Step 5: Sleep in milliseconds
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("I'm awake now!")

	// Step 6: Time taken
	t0 := time.Now()
	time.Sleep(2000 * time.Millisecond)
	t1 := time.Now()
	fmt.Printf("The sleep took approx %v to run.\n", t1.Sub(t0))

}
