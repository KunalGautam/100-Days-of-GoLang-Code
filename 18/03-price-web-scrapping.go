package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	// get response
	resp, err := http.Get("https://www.amazon.in/How-Win-Friends-Influence-People/dp/8129140179/")
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
	data := string(body)
	// Regex Query in progress
	regexQuery := regexp.MustCompile("")

	price := regexQuery.FindStringSubmatch(data)
	fmt.Println("Price: ", price)
}
