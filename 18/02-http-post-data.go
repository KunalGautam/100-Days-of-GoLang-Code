package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Can be better way done by defining struct, but due to time limitation everything is defined as string
	data, _ := json.Marshal(map[string]string{
		"Id":       "78912",
		"Customer": "Jason Sweet",
		"Quantity": "1",
		"Price":    "18.00",
	})
	dataToSend := bytes.NewBuffer(data)
	resp, err := http.Post("https://reqbin.com/echo/post/json", "application/json", dataToSend)
	if err != nil {
		panic(err)
	}
	resp.Header.Set("Accept", "application/json")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
