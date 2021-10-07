package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

type indexData struct {
	// Capitalize First letter of variable to make it available globally, else will have error of is an unexported field of struct type. The renderer package use the reflect package in order to get/set fields values the reflect package can only access public/exported struct fields.
	Title    string
	User     string
	Password string
}

func main() {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Panicln(err)
	}
	t2, err := template.ParseFiles("template/index2.html")
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

	data = indexData{
		Title: "Title of the Page2",
		User:  "User2",
		// Enable or disable if required.
		//Password: "p455w0rd",
	}
	err = t2.Execute(buffer, data)
	if err != nil {
		log.Panicln(err)
	}
	renderedOutput = buffer.String()
	fmt.Println(renderedOutput)
}
