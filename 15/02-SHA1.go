package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	data := "Hello World!"

	s := sha1.New()
	encrypt, _ := s.Write([]byte(data))
	bs := s.Sum(nil)
	fmt.Println(encrypt)
	fmt.Printf("%x \n", bs)
}
