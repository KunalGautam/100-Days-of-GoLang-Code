package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

type IndexData struct {
	UserName   string
	Department string
}
type pageMeta struct {
	Title       string
	WelcomeText string
}

type data struct {
	Meta pageMeta
	Data []IndexData
}

func main() {
	t, err := template.ParseFiles("tpl/index.tpl")
	if err != nil {
		log.Panicln(err)
	}

	metaData := pageMeta{
		Title:       "Page Title",
		WelcomeText: "Hello World!",
	}
	dataOfIndex := []IndexData{
		{
			"User1",
			"Department1",
		}, {
			"User2",
			"Department2",
		}, {
			"User3",
			"Department3",
		}, {
			"User4",
			"Department4",
		},
	}
	data := data{
		Meta: metaData,
		Data: dataOfIndex,
	}
	buffer := &bytes.Buffer{}
	err = t.Execute(buffer, data)
	if err != nil {
		log.Panicln(err)
	}
	renderedOutput := buffer.String()
	fmt.Println(renderedOutput)

}
