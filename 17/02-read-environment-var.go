package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("User:", os.Getenv("USER"))
}
