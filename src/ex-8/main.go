package main

// html dosyasını okuyup localhost:8080/ 'de gosterecegiz
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//dosyayı alalım
func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var body, _ = loadFile("page.html")
	fmt.Fprintf(w, body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
