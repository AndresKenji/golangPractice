package server

import (
	vmware "consoleapp/internal/vcenter"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func vmMksHandler(w http.ResponseWriter, r *http.Request) {
	VmName := r.URL.Query().Get("vmname")
	VcName := r.URL.Query().Get("vcenter")
	if VmName == "" || VcName == "" {
		http.Error(w, "vmname and vcenter are required", http.StatusBadRequest)
		return
	}

	log.Println("Iniciando sesión en:", VcName)
	vcenter, err := vmware.NewVcenter(VcName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ticket, err := vcenter.GetWmksParams(VmName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	out, err := json.Marshal(WmksResponse{
		Ticket: ticket.Ticket,
		Host: ticket.Host,
		VmName: VmName,
		Port: ticket.Port,
		URL: ticket.Url,
		Vcenter: VcName,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func consoleProxyWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// VmName := r.URL.Query().Get("vmname")
	VcName := r.URL.Query().Get("vcenter")
	vcHost := r.URL.Query().Get("host")
	token := r.URL.Query().Get("token")


	// if VmName == "" || VcName == "" {
	// 	http.Error(w, "vmname and vcenter are required", http.StatusBadRequest)
	// 	return
	// }
	if vcHost == "" || token == "" || VcName == "" {
		http.Error(w, "host, token and vcenter are required", http.StatusBadRequest)
		return
	}

	// log.Println("Iniciando sesión en:", VcName)
	// vcenter, err := vmware.NewVcenter(VcName)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// targetURL, err := vcenter.GetWmksTicket(VmName)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	targetURL := fmt.Sprintf("wss://%s/ticket/%s",vcHost,token)

	headers := http.Header{
		"User-Agent":             r.Header["User-Agent"],
		"Cookie":                 r.Header["Cookie"],
		"Sec-WebSocket-Protocol": []string{"binary"},
	}

	if auth := r.Header.Get("Authorization"); auth != "" {
		headers.Set("Authorization", auth)
	}
	headers.Set("Pragma", "no-cache")
	headers.Set("Cache-Control", "no-cache")
	headers.Set("Host", VcName)
	headers.Set("Origin", fmt.Sprintf("https://%s", VcName))

	tlsConfig := &tls.Config{
		MinVersion:         tls.VersionTLS12,
		CipherSuites:       []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384},
		ServerName:         VcName,
		InsecureSkipVerify: true,
	}

	wsDialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		EnableCompression: true,
		HandshakeTimeout:  25 * time.Second,
		TLSClientConfig: tlsConfig,
	}

	connBackend, response, err := wsDialer.Dial(targetURL, headers)
	if err != nil {
		log.Println("Error al conectar con el backend:", err)
		if response != nil {
			body, _ := io.ReadAll(response.Body)
			log.Println("Respuesta del backend:", string(body))
			response.Body.Close()
		}
		http.Error(w, "Error en handshake con backend", http.StatusBadGateway)
		return
	}
	defer connBackend.Close()

	// Actualizar la conexión del cliente
	upgrader := websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		Subprotocols:    []string{"binary"},
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
	}
	connClient, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar conexión con el cliente:", err)
		return
	}
	defer connClient.Close()

	log.Println("Iniciando streaming a:",targetURL)

	// Proxy bidireccional
	go func() {
		defer connClient.Close()
		defer connBackend.Close()
		for {
			mt, msg, err := connClient.ReadMessage()
			if err != nil {
				return
			}
			if err := connBackend.WriteMessage(mt, msg); err != nil {
				return
			}
		}
	}()

	for {
		mt, msg, err := connBackend.ReadMessage()
		if err != nil {
			return
		}
		if err := connClient.WriteMessage(mt, msg); err != nil {
			return
		}
	}
}
