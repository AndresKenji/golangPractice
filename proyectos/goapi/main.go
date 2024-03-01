package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"kenji.goapi/goapi/routes"
)




func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8800",router)

}
