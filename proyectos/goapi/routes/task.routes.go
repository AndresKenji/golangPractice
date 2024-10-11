package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"kenji.goapi/goapi/db"
	"kenji.goapi/goapi/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	id := mux.Vars(r)["id"]
	db.DB.First(&task, id)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	id := mux.Vars(r)["id"]
	db.DB.First(&task, id)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) // 404
}
