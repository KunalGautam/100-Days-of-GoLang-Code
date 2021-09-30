package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// get response
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status Code: ", resp.StatusCode)
	fmt.Println("Protocol: ", resp.Proto)
	fmt.Println("Header: ", resp.Header, "\n\n\n\n")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
