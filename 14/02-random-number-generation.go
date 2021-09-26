package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	num := rand.NewSource(time.Now().Unix())
	fmt.Println(rand.New(num).Intn(1000))
}
