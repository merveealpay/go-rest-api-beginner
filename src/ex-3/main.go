package main

import (
	"fmt"
	"log"
	"net/http"
)

//custom handler
func main() {
	mux := http.NewServeMux()

	x1 := &messageHandler{"Bu bir mesajdır."}
	x2 := &messageHandler{"Bu da bir mesajdır."}

	mux.Handle("/bir", x1)
	mux.Handle("/iki", x2)

	log.Println("Listening")

	http.ListenAndServe(":8080", mux)

}

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, x.message) //anlık olarak pointer uzerınden gelen veri
}
