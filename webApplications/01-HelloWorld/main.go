package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello, World")
		fmt.Println("Bytes written:",n)
		if err != nil {

		}
	})
}