package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	//serverURL := "ws://localhost:8801/ws"
	serverURL := "wss://cloudconsole.ifxcorp.com:9500/bogp/wbpdlx28/5760924002b79320"
	headers := http.Header{}
	headers.Set("Host", "cloudconsole.ifxcorp.com:9500")
	conn, response, err := websocket.DefaultDialer.Dial(serverURL, headers)
	if err != nil {
		body, _ := io.ReadAll(response.Body)
		log.Println(string(body))
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