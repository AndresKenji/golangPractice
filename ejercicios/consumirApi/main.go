package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendTeamsReport(htmlReport string, identifiers []string) (*http.Response, error) {
	url := "https://nfgw.ifxcorp.com/api/Msg/TeamsSendChatMessage"
	// Crear un mapa que contiene el contenido del mensaje a enviar
	data := map[string]interface{}{
		"message":     htmlReport,
		"identifiers": identifiers,
	}

	// Convertir el mapa de datos a JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Crear una solicitud HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Agregar los encabezados a la solicitud HTTP
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/plain")

	// Realizar la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)

	// Retornar la respuesta HTTP de la solicitud
	return resp, nil
}

func main() {

	identifiers := []string{"19:38a33eec37364c4a8cb26725630698a1@thread.v2"}
	resp, err := SendTeamsReport("<html><body><h1>Hello, Teams!</h1></body></html>", identifiers)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
