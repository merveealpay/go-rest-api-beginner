package main

import (
	"fmt"
	"net/http"
)

var data string

const (
	// DataIndex is ...
	DataIndex string = "Index Page"
	// DataAbout is ...
	DataAbout string = "About Page"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pathVar := r.URL.Path[1:]

	if len(pathVar) == 0 {
		data = DataIndex
	} else {
		data = fmt.Sprintf("Path: %s", pathVar)
	}

	w.Write([]byte(data))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(DataAbout))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)
}
