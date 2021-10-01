package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	openBrowser("http://localhost:8081")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q \nTo Exit Program  go to http://localhost:8081/exit\n\nInteresting URL: http://localhost:8081/hi", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Awesome!")
	})

	http.HandleFunc("/exit", func(_ http.ResponseWriter, r *http.Request) {
		os.Exit(11)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
