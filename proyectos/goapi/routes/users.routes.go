package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"kenji.goapi/goapi/db"
	"kenji.goapi/goapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
	
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user models.User
	db.DB.First(&user,id)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode((&user))

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	} else {

	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user models.User
	db.DB.First(&user,id)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("User not found"))
		return
	} else {
		db.DB.Unscoped().Delete(&user)
		w.WriteHeader(http.StatusOK) // 200
		w.Write([]byte("Deleted"))
	}
}
