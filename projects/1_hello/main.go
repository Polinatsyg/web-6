package main

// некоторые импорты нужны для проверки
import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello,web!"))
}
func main() {
	http.HandleFunc("/get", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}


