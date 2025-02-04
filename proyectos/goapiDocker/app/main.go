package main

import (
	"apialertservice/ifxcorp.com/api"
	"fmt"
	"log"
)

func main() {
	fmt.Println("iniciando Api")

	server := api.NewApiServer(":8801")
	err := server.Run()
	if err != nil {
		log.Panicln(err)
	}
	defer fmt.Println("Terminando ejecución")

}
