package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Index Page"))
	//w.WriteHeader(http.StatusOK)

	//burada localhost:8080'den sonra yazdıgımız degeri pathVar' olarak aldım
	pathVar := r.URL.Path[1:]
	data := ""

	if len(pathVar) > 0 {
		data = "Path: " + pathVar
	} else {
		data = "Index Page"
	}

	w.Write([]byte(data))

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func main() {

	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)

}
