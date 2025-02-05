package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Permitir todas las conexiones (en producción, restringe esto)
		return true
	},
}

func wsProxyHandler(targetURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Actualizar la conexión del cliente a WebSocket
		clientConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error al actualizar la conexión del cliente:", err)
			return
		}
		defer clientConn.Close()

		// 2. Conectar al servidor WebSocket de destino
		dstURL, err := url.Parse(targetURL)
		if err != nil {
			log.Println("Error al parsear la URL de destino:", err)
			return
		}

		dstConn, _, err := websocket.DefaultDialer.Dial(dstURL.String(), nil)
		if err != nil {
			log.Println("Error al conectar con el servidor de destino:", err)
			return
		}
		defer dstConn.Close()

		// 3. Canal para enviar mensajes del cliente al servidor de destino
		go func() {
			for {
				messageType, message, err := clientConn.ReadMessage()
				if err != nil {
					log.Println("Error leyendo del cliente:", err)
					break
				}
				if err := dstConn.WriteMessage(messageType, message); err != nil {
					log.Println("Error escribiendo al servidor de destino:", err)
					break
				}
			}
		}()

		// 4. Canal para enviar mensajes del servidor de destino al cliente
		for {
			messageType, message, err := dstConn.ReadMessage()
			if err != nil {
				log.Println("Error leyendo del servidor de destino:", err)
				break
			}
			if err := clientConn.WriteMessage(messageType, message); err != nil {
				log.Println("Error escribiendo al cliente:", err)
				break
			}
		}
	}
}

func main() {
	// Configurar el endpoint del proxy
	http.HandleFunc("/ws", wsProxyHandler("ws://localhost:8080/ws"))

	// Iniciar el servidor
	log.Println("Servidor proxy WebSocket iniciado en :8801")
	if err := http.ListenAndServe(":8801", nil); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}

