package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	serverURL := "ws://localhost:8801/ws"
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatal("Error al conectar con el servidor WebSocket:", err)
	}
	defer conn.Close()

	log.Println("Conectado al servidor WebSocket")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Conexión cerrada:", err)
			break
		}
		log.Println("Mensaje recibido:", string(msg))
	}
}