package main

import (
	"log"
	"net/http"
)

func main() {
	//mux := http.NewServeMux()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola, mundo!"))
	})

	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}