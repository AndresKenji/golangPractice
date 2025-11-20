package main

import (
	"fmt"
)

func main() {

	fmt.Println("Rest API Go 1.22")
	// mux := http.NewServeMux()

	// mux.HandleFunc("/",func (w http.ResponseWriter, r *http.Request)  {
	// 	fmt.Fprintf(w,"Hello world!")
	// })

	// mux.HandleFunc("GET /comment",func (w http.ResponseWriter, r *http.Request)  {
	// 	fmt.Fprintf(w,"list of all comments")
	// })

	// mux.HandleFunc("GET /comment/{id}",func (w http.ResponseWriter, r *http.Request)  {
	// 	id := r.PathValue("id")

	// 	fmt.Fprintf(w,"return a single comment with id = %v", id)
	// })

	// mux.HandleFunc("POST /comment",func (w http.ResponseWriter, r *http.Request)  {
	// 	fmt.Fprintf(w,"Post a new comment")
	// })

	// if err := http.ListenAndServe("localhost:8801",mux); err != nil {
	// 	fmt.Println(err.Error())
	// }

	server := NewAPIserver("localhost:8801")
	server.Run()
}
