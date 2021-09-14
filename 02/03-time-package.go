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
}
