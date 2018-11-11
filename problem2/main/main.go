package main

import (
	"fmt"
	"gophercises/problem2"
	"net/http"
	"io/ioutil"
	"flag"
)

var (
	yamlFilePath string
	jsonFilePath string
)

func init() {
	flag.StringVar(&yamlFilePath, "yamlFilePath", "main/sampleurl.yaml", "Yaml file from which to read url's")
	flag.StringVar(&jsonFilePath, "jsonFilePath", "main/url.json", "JSON file from which to read url's")
	flag.Parse()
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := problem2.MapHandler(pathsToUrls, mux)

	 //Build the YAMLHandler using the mapHandler as the
	 //fallback
	yaml, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := problem2.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	json, err := ioutil.ReadFile(jsonFilePath)
	jsonHandler, err := problem2.JSONHandler([]byte (json), yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
