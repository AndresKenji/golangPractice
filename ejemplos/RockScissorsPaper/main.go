package main

import (
	"log"
	"net/http"
	"kenji.rsp/rsp/handlers"
)




func main() {
	// Objeto router para recibir rutas
	router := http.NewServeMux()
	// Manejo de archivos estaticos 
	fs := http.FileServer(http.Dir("static"))
	// Registrar la ruta para los archivos estaticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	// Agregar rutas al router
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/about", handlers.About)
	port := ":8801"
	log.Printf("Servidor escuchando en http://localhost:%v\n",port)
	log.Fatal(http.ListenAndServe(port,router))
	
}