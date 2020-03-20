package main

import (
	"io"
	"net/http"
)

//mux, bir kere olustur bunun uzerınden handle et.
func main() {

	var i ironman
	var w wolverine

	mux := http.NewServeMux()
	mux.Handle("/ironman", i)
	mux.Handle("/wolverine", w)
	http.ListenAndServe(":8080", mux)
}

type ironman int

func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr.Iron!")
}

type wolverine int

func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr.Wolverine!")
}
