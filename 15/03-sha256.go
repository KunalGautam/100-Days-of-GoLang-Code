package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := "password"

	s := sha256.New()
	encrypt, _ := s.Write([]byte(data))
	bs := s.Sum(nil)
	fmt.Println(encrypt)
	fmt.Printf("%x \n", bs)

}
