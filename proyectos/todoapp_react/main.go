package main

import (
	"log"
	"todoapp/api"
	"todoapp/db"
	"todoapp/models"
)

func main() {
	// conectar a la db
	db.DBConnection()
	// migrar modelos a la db
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})


	// Creando el servidor para la Api
	server := api.NewApiServer("localhost:8801")
	err := 	server.Run(); if err != nil {
		log.Panicln(err)
	}

}