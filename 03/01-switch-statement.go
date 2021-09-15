package main

import (
	"fmt"
	"time"
)

func main() {
	// Step1
	a := 3
	switch a % 2 {
	case 0:
		fmt.Println("Number is Even")
	case 1:
		fmt.Println("Number is Odd")
	}

	// Step 2: Set default
	switch a % 5 {
	case 0:
		fmt.Println("Number is divisible by 5")
	default:
		fmt.Println("Number is not divisible by 5")

	}

	// Step 3: weekday or weekend?
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Step 4: Don't pass expression to switch

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

}
