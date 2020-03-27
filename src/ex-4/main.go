package main

import (
	"fmt"
	"net/http"
)

//merhaba + path degeri
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)

	fmt.Println("Web Server")
}
