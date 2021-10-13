package main

import (
	"fmt"
	"gopkg.in/yaml.v2"

	//yaml "gopkg.in/yaml.v2"
	"net/http"
	"os"
)

func main() {
	yamlData, err := os.ReadFile("urls.yaml")
	if err != nil {
		panic(err)
	}

	var pathsToUrls []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}
	if err := yaml.Unmarshal(yamlData, &pathsToUrls); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	for _, url := range pathsToUrls {
		uri := url.Path
		link := url.URL
		mux.HandleFunc(uri, func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, link, http.StatusFound)
		})
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello")
		return
	})

	http.ListenAndServe(":8080", mux)

}
