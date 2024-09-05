package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"hwmonitor/internal/hardware"
	"hwmonitor/internal/server"
)

func main() {
	fmt.Println("Starting system monitor...")
	srv := server.NewServer()
	go func(s *server.Server) {
		for {
			systemSection, err := hardware.GetSystemSection()
			if err != nil {
				fmt.Println("Error fetching system section:", err)
				continue
			}

			diskSection, err := hardware.GetDiskSection()
			if err != nil {
				fmt.Println("Error fetching disk section:", err)
				continue
			}

			cpuSection, err := hardware.GetCpuSection()
			if err != nil {
				fmt.Println("Error fetching CPU section:", err)
				continue
			}

			timeStamp := time.Now().Format("2006-01-02 15:04:05")

			html := `
			<strong hx-swap-oob="innerHTML:#update-timestamp">Last Update ` + timeStamp + `</strong>
			<div hx-swap-oob="innerHTML:#system-value">`+systemSection.GetHtml()+`</div>
			<div hx-swap-oob="innerHTML:#disk-value">`+diskSection.GetHtml()+`</div>
			<div hx-swap-oob="innerHTML:#cpu-value">`+cpuSection.GetHtml()+`</div>
			`
			s.Broadcast([]byte(html))

			time.Sleep(3 * time.Second)
		}
	}(srv)

	err := http.ListenAndServe(":8080", &srv.Mux) 
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
