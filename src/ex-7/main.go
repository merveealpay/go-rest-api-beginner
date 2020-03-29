package main

// bu ornekte, localhost:8080/api/user/5 ---->>> 5 id'li user gelecek.

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux" //mux ın fonksiyonlarına erişmek için, gorilla webtoolkit
)

type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r) //requestin variable'larını getir
	id := urlParams["id"]
	messageConcat := "Kullanici ID: " + id

	message := API{messageConcat}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Hata olustu.")
	}
	fmt.Fprintf(w, string(output))

}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/user/{id}", Hello)
	http.Handle("/", gorillaRoute)

	http.ListenAndServe(":8080", nil)

}
