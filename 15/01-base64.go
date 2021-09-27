package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hello World!"

	encode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encode)
	decode, _ := base64.StdEncoding.DecodeString(string(encode))
	fmt.Println(string(decode))
}
