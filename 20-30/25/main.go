package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

type indexData struct {
	// Capitalize First letter of variable to make it available globally, else will have error of is an unexported field of struct type
	Title string
	User  string
}

func main() {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Panicln(err)
	}
	data := indexData{
		Title: "Title of the Page",
		User:  "User1",
	}
	buffer := &bytes.Buffer{}
	err = t.Execute(buffer, data)
	if err != nil {
		log.Panicln(err)
	}
	renderedOutput := buffer.String()
	fmt.Println(renderedOutput)
}
