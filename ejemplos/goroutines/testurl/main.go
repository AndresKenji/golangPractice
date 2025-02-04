package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// canal := make(chan int) // Asi se declara
	// canal <- 15             // Asi se envian datos al canal
	// valor := <-canal 		// Asi se reciben datos del canal
	ch := make(chan string)

	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	//time.Sleep(5 * time.Second)

	elapsed := time.Since(start)

	fmt.Printf("¡Listo! ¡Tomó %v segundos!\n", elapsed.Seconds())

}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: ¡%s está abajo!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: ¡%s está en funcionamiento!\n", api)

}
