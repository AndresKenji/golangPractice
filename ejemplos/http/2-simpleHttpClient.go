package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	req, err := http.NewRequest(http.MethodGet, "https://pokeapi.co/api/v2/pokemon/pikachu", nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Codigo de respuesta:", resp.StatusCode)
	fmt.Println("Content-Type", resp.Header["Content-Type"])
	cuerpo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("------------")
	fmt.Println(string(cuerpo))

}
