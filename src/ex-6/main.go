package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}

func main() {

	apiRoot := "/api"

	//.../api
	//http://localhost:8080/api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API HOME"}
		output, err := json.Marshal(message)
		checkError(err)
		//w.Header().Set("Content-Type","application/json")
		fmt.Fprintf(w, string(output))

	})

	//.../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FirstName: "Merve", LastName: "Alpay", Age: 23},
			User{ID: 2, FirstName: "Kumsal", LastName: "Deniz", Age: 24},
			User{ID: 3, FirstName: "Gunes", LastName: "Bulut", Age: 25},
		}
		message := users
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	//.../api/me

	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{0, "adminMerve", "adminAlpay", 23}
		message := user
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)

}

//hata kontrolu yapalim
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1) //uygulamadan cik
	}
}
