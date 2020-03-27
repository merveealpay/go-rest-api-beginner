package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func (human Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	human.FirstName = "Merve"
	human.LastName = "Alpay"
	human.Age = 23

	r.ParseForm()

	//sunucudan form bilgisini almak icin

	fmt.Println(r.Form)

	//url'in path bilgisini almak icin

	fmt.Println("path", r.URL.Path[1:])

	//html
	fmt.Fprint(w, "<table><tr><td><b>Ad</b></td><td><b>Soyad</b></td><td><b>Yaş</b></td></tr><tr><td>"+human.FirstName+"</td><td>"+human.LastName+"</td><td>"+strconv.Itoa(human.Age)+"</td></tr><tr></tr><tr></tr><tr><td><b>Path</b></td><td>"+r.URL.Path+"</td></tr></table>")

}

func main() {

	var human Human
	err := http.ListenAndServe(":9000", human)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
