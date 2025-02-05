package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar la conexión:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	log.Println("Nuevo cliente conectado")

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Cliente desconectado")
			delete(clients, conn)
			break
		}
	}
}

func broadcastMessages() {
	for {
		message := fmt.Sprintf("Mensaje en %v", time.Now().Format(time.RFC3339))
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Error al enviar mensaje:", err)
				client.Close()
				delete(clients, client)
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go broadcastMessages()

	log.Println("Servidor WebSocket en ejecución en ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}