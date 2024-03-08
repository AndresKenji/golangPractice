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
	// User Routes
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Tasks Routes
	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.PostTasksHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8800",router)

}
