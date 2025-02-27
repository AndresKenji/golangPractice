package server

import (
	"net/http"
)

func (c *AppConfig) RegisterRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server Listening"))
	})
	router.HandleFunc("/consoleproxy/", consoleProxyWebSocketHandler)
	router.HandleFunc("/vmmks/", vmMksHandler)

	return router
}

type WmksResponse struct {
	Ticket string `json:"ticket"`
	Host string `json:"host"`
	VmName string `json:"vmname"`
	Port int32 `json:"port"`
	URL string `json:"url"`
	Vcenter  string`json:"vcenter"`
}