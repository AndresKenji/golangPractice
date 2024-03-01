package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"kenji.goapi/goapi/db"
	"kenji.goapi/goapi/models"
	"kenji.goapi/goapi/routes"
)




func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8800",router)

}
