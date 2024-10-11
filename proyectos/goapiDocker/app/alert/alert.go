package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SlackMsg struct {
	Text  string `json:"text"`
	Error bool   `json:"err"`
}

func SlackNotification(msg *SlackMsg) (*http.Response, error) {
	slackError := ""
	slackWarning := ""

	payload := struct {
		Text string `json:"text"`
	}{
		Text: fmt.Sprintf("[cron-inconsistencies]--%s-- %s", time.Now().Format("2006-01-02 15:04:05"), msg.Text),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := slackError
	if !msg.Error {
		url = slackWarning
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type TeamsMsg struct {
	Text        string   `json:"text"`
	Identifiers []string `json:"identifiers"`
}

func SendTeamsReport(msg *TeamsMsg) (*http.Response, error) {
	url := ""
	// Crear un mapa que contiene el contenido del mensaje a enviar
	data := map[string]interface{}{
		"message":     msg.Text,
		"identifiers": msg.Identifiers,
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

	// Retornar la respuesta HTTP de la solicitud
	return resp, nil
}
