package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http" // пакет для поддержки HTTP протокола
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                   // Разрешить доступ всем источникам (можно указать конкретный домен вместо "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Разрешённые методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Разрешённые заголовки

	// Обрабатываем OPTIONS-запросы
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello,%s!", name)
}
func main() {
	http.HandleFunc("/api/user", handler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
