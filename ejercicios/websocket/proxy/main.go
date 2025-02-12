package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var wsDialer = websocket.Dialer{
	Proxy:           http.ProxyFromEnvironment,
	TLSClientConfig: &tls.Config{
		//InsecureSkipVerify: true,
	},
	EnableCompression: true,
	HandshakeTimeout: 25 * time.Second,
}

func loadTLSConfig(certFile, keyFile, caFile string) (*tls.Config, error) {
    cert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        return nil, err
    }

    config := &tls.Config{
        Certificates:       []tls.Certificate{cert},
        MinVersion:         tls.VersionTLS12,
        CipherSuites:       []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384},
        ServerName:         "cloudconsole.ifxcorp.com",
        InsecureSkipVerify: true,  // Temporal para pruebas
    }

    if caFile != "" {
        caCert, _ := os.ReadFile(caFile)
        caCertPool := x509.NewCertPool()
        caCertPool.AppendCertsFromPEM(caCert)
        config.RootCAs = caCertPool
        config.InsecureSkipVerify = false  // Solo si el CA es válido
    }

    return config, nil
}

func wsProxy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	// Construye la URL correcta sin slash adicional
	targetURL := "wss://cloudconsole.ifxcorp.com:9500" + r.URL.Path
	log.Println("Conectando a:", targetURL)

	// Copia las cabeceras del cliente
	headers := http.Header{
		"User-Agent":          r.Header["User-Agent"],
		"Origin":              r.Header["Origin"],
		"Cookie":              r.Header["Cookie"],
		"Sec-WebSocket-Protocol": []string{"binary"},
		"Host":                []string{"cloudconsole.ifxcorp.com:9500"},  // <--- Nuevo
	}
	
	// Copia headers de autenticación si existen
	if auth := r.Header.Get("Authorization"); auth != "" {
		headers.Set("Authorization", auth)
	}
	headers.Set("Pragma", "no-cache")
	headers.Set("Cache-Control", "no-cache")
	headers.Set("Host", "cloudconsole.ifxcorp.com:9500")  // Necesario para enrutamiento
	
	// Cargar el certificado y la clave para la conexión TLS
	tlsConfig, err := loadTLSConfig("./ifxcorp.crt", "./ifxcorp.key", "")
	if err != nil {
		log.Println("Error al cargar certificados:", err)
		http.Error(w, "Error al cargar certificados", http.StatusInternalServerError)
		return
	}

	// Usar la configuración TLS cargada en el Dialer
	wsDialer.TLSClientConfig = tlsConfig

	// Conectar al backend con las cabeceras
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
		Subprotocols:    []string{"binary"},  // <--- Clave para VMware
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	connClient, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar conexión con el cliente:", err)
		return
	}
	defer connClient.Close()

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

func main() {
    log.Println("Iniciando proxy...")
    http.HandleFunc("/", wsProxy)
    server := &http.Server{
        Addr: ":8080",
        ReadTimeout: 15 * time.Second,
        WriteTimeout: 15 * time.Second,
    }
    log.Fatal(server.ListenAndServe())
}
