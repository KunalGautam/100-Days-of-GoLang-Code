package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var count = 0

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	countCookie, err := req.Cookie("counter")

	if err != nil {
		count = 1
		setCookie(count, res, req)
	} else {
		count, err = strconv.Atoi(countCookie.Value)
		if err != nil {
			fmt.Fprintln(res, "ERROR!")
		}
		count++
		setCookie(count, res, req)
	}
	fmt.Fprintln(res, "YOUR COUNTER:", count)

}

func setCookie(i int, w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "counter",
		Value: strconv.Itoa(i),
	})
}
