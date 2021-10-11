package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", getMethod)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func getMethod(w http.ResponseWriter, req *http.Request) {
	getData := req.FormValue("name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="GET">
	 <input type="text" name="name" placeholder="Enter Your Name">
	 <input type="submit">
	</form>
	<br>`+getData)
}
