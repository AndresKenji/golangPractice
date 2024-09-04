package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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

			<td hx-swap-oob="innerHTML:#hostname-value">` + systemSection.Hostname + `</td>

			<td hx-swap-oob="innerHTML:#memory-value">` + strconv.Itoa(systemSection.MemoryMB) + `</td>

			<td hx-swap-oob="innerHTML:#free-memory-value">` + strconv.Itoa(systemSection.FreeMemoryMB) + `</td>

			<td hx-swap-oob="innerHTML:#os-value">` + systemSection.Os + `</td>

			<td hx-swap-oob="innerHTML:#platform-value">` + systemSection.Platform + `</td>

			<td hx-swap-oob="innerHTML:#uptime-value">` + strconv.FormatUint(uint64(systemSection.Uptime), 10) + `</td>

			<td hx-swap-oob="innerHTML:#disk-total-value">` + strconv.Itoa(int(diskSection.TotalSpaceGB)) + `</td>

			<td hx-swap-oob="innerHTML:#disk-used-value">` + strconv.Itoa(int(diskSection.UsedSpaceGB)) + `</td>

			<td hx-swap-oob="innerHTML:#disk-free-value">` + strconv.Itoa(int(diskSection.FreeSpaceGB)) + `</td>

			<td hx-swap-oob="innerHTML:#cpu-value">` + cpuSection.CPU + `</td>

			<td hx-swap-oob="innerHTML:#cpu-cores-value">` + strconv.Itoa(cpuSection.Cores) + `</td>

			<td hx-swap-oob="innerHTML:#cpu-logicalcores-value">` + strconv.Itoa(cpuSection.LogicalCores) + `</td>
			
			<td hx-swap-oob="innerHTML:#cpu-percent-value">` + strconv.Itoa(int(cpuSection.Percent)) + `%</td>
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
