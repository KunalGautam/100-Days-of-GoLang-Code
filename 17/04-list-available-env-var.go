package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Environ() list all available env vars and the values
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], "\t\t\t", pair[1])
	}
}
