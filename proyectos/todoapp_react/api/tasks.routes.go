package api

import (
	"encoding/json"
	"log"
	"net/http"
	"todoapp/db"
	"todoapp/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	userid, err := GetCookieUSer(r)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
	}
	db.DB.Where("user_id = ?", userid).Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	userid, err := GetCookieUSer(r)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
	}
	var task models.Task
	id := r.PathValue("id")
	db.DB.First(&task, id)
	if task.ID == 0 || task.UserId != uint(userid) {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func ChangeTaskStateHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	id := r.PathValue("id")
	db.DB.First(&task, id)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}
	task.Done = !task.Done
	db.DB.Save(&task)
	json.NewEncoder(w).Encode(&task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	id := r.PathValue("id")
	db.DB.First(&task, id)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}

	// Crear un decodificador para leer el cuerpo del request
	decoder := json.NewDecoder(r.Body)

	// Definir una estructura o variable para almacenar los datos del cuerpo del request
	var datos models.Task
	// Decodificar el cuerpo del request en la estructura o variable definida
	err := decoder.Decode(&datos)
	if err != nil {
		// Manejar el error si ocurre al decodificar el cuerpo del request
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al decodificar el cuerpo del request"))
		return
	}
	task.Description = datos.Description
	task.Title = datos.Title
	db.DB.Save(&task)
	json.NewEncoder(w).Encode(&task)
}

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	userid, err := GetCookieUSer(r)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
	}
	log.Println("User Id: ", userid)
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	task.UserId = uint(userid)
	createdTask := db.DB.Create(&task)
	err = createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	id := r.PathValue("id")
	db.DB.First(&task, id)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) // 404
}
