package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("SOME_ENV_VAR1", "1")
	fmt.Println("SOME_ENV_VAR1:", os.Getenv("SOME_ENV_VAR1"))
	os.Unsetenv("SOME_ENV_VAR1")
}
