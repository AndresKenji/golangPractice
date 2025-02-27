package main

import (
	"consoleapp/internal/server"
	"log"
)

func main() {

	config := server.AppConfig{
		Port:         80,
		ReadTimeout:  10,
		WriteTimeout: 10,
	}

	srv := server.NewServer(&config)
	log.Println("Server running on port:", config.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err.Error())
	}

}
