package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", postSubmission)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func postSubmission(w http.ResponseWriter, req *http.Request) {
	postData := req.FormValue("name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="name" placeholder="Enter Your name">
	 <input type="submit">
	</form>
	<br>`+postData)
}
