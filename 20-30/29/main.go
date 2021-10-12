package main

import (
	yaml "gopkg.in/yaml.v2"
	"net/http"
	"os"
)

func main() {
	readYamlFile, err := os.ReadFile("urls.yaml")
	if err != nil {
		panic(err)
	}

	handler, err := YAMLHandler(readYamlFile)

	http.ListenAndServe(":8080", handler)

}

func YAMLHandler(yamldata []byte) (http.HandlerFunc, error) {
	var fallback http.Handler
	var pathsToUrls []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}
	if err := yaml.Unmarshal(yamldata, &pathsToUrls); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathtourl := range pathsToUrls {
			if pathtourl.Path == r.URL.Path {
				http.Redirect(w, r, pathtourl.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}
