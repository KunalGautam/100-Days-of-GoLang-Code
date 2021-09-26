package main

import (
	"fmt"
	"net/url"
)

func main() {
	urltoparse := "http://user%20name:passw0rd@example.com:8080/path?q=test"
	u, err := url.Parse(urltoparse)
	if err != nil {
		panic(err)
	}

	// Prints protocol
	fmt.Println(u.Scheme)

	// Print Auth details
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	// Print URI
	fmt.Println(u.Host)

	// print port
	fmt.Println(u.Port())

	// print path
	fmt.Println(u.RequestURI())
	fmt.Println(u.Path)

	// print query
	fmt.Println(u.Query())

}
