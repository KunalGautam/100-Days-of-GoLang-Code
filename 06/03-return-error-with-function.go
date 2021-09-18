package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	} else {
		// If not an error, return nil
		return math.Sqrt(f), nil
	}
}

func main() {
	val, err := Sqrt(-4.0)
	if err != nil {
		// Create panic and exit the prog
		log.Panicln(err)
	}
	fmt.Println(val)
}
