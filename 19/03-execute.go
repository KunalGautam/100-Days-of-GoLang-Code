package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	lsPath, err := exec.LookPath("ls")
	if err != nil {
		log.Panic(err)
	}
	// Path of binary
	fmt.Println(lsPath)
	arguments := []string{"ls", "-l", "-R", "-a", "-h"}
	env := os.Environ()
	err = syscall.Exec(lsPath, arguments, env)
	if err != nil {
		log.Panic(err)
	}
}
