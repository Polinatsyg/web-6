package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"log"
	"net/http"
	"strconv" // вдруг понадобиться вам ;)
)

var counter = 0

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		a, err := strconv.Atoi(r.FormValue("count"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}
		counter += a
		w.Write([]byte("OK!"))
		return
	} else if r.Method == "GET" {
		w.Write([]byte(strconv.Itoa(counter)))
		return
	}
	w.Write([]byte("Разрешен только метод POST и GET!"))
}
func main() {
	http.HandleFunc("/count", handler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
