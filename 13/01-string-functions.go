package main

import (
	"fmt"
	//https://pkg.go.dev/strings
	"strings"
)

func main() {

	fmt.Println("Contains:  	", strings.Contains("test", "es"))
	fmt.Println("Count:     	", strings.Count("test", "t"))
	fmt.Println("HasPrefix: 	", strings.HasPrefix("test", "te"))
	fmt.Println("Join:      	", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    	", strings.Repeat("a", 5))
	fmt.Println("Replace:   	", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   	", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     	", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   	", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   	", strings.ToUpper("test"))
	fmt.Println("First Uppercase:", strings.Title(strings.ToLower("HELLO")))
	fmt.Println("First Uppercase:", strings.Title(strings.ToLower(strings.ToUpper("hola voila"))))
	fmt.Println("Len: 		", len("hello"))
	fmt.Println("Char:		", "hello"[1])
}
