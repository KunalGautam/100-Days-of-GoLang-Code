package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", fileUploadMethod)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func fileUploadMethod(w http.ResponseWriter, req *http.Request) {

	var s string
	fmt.Println("Request Method", req.Method)
	if req.Method == http.MethodPost {

		// open
		f, _, err := req.FormFile("fileupload1")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="fileupload1">
	<input type="submit">
	</form>
	<br><pre>`+s+`</pre>`)
}
