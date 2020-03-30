package main

import (
	"fmt"
	"io"
	"net/http"
)

type hero struct {
	ID          int
	Name, Quote string
}

var (
	ironman = hero{
		ID:    1,
		Name:  "ironman",
		Quote: "Sometimes you gotta run before you can walk.",
	}
	wolverine = hero{
		ID:    1,
		Name:  "wolverine",
		Quote: "Nature made me a freak. Man made me a weapon. And god made it last too long.",
	}
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/ironman", ironman)
	mux.Handle("/wolverine", wolverine)

	http.ListenAndServe(":8080", mux)
}

func (hero hero) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, fmt.Sprintf("%s : %s", hero.Name, hero.Quote))
}
