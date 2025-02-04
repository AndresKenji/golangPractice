package api

import (
	"apialertservice/ifxcorp.com/alert"
	"encoding/json"
	"net/http"
)

func SendTeamsMsgHandler(w http.ResponseWriter, r *http.Request) {
	var body alert.TeamsMsg
	json.NewDecoder(r.Body).Decode(&body)
	response, err := alert.SendTeamsReport(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if response.Status == "200 OK" {
		w.Write([]byte("Alerta enviada con exito"))
	}

}

func SendSlackMsgHandler(w http.ResponseWriter, r *http.Request) {
	var body alert.SlackMsg
	json.NewDecoder(r.Body).Decode(&body)
	response, err := alert.SlackNotification(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if response.Status == "200 OK" {
		w.Write([]byte("Alerta enviada con exito"))
	}

}
