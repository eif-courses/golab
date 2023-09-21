package controllers

import (
	"encoding/json"
	"github.com/eif-courses/golab/services"
	"github.com/eif-courses/golab/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

var user services.User

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	all, err := user.GetAllUsers()
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}
	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"users": all})
	utils.ServerErrorHTTP(err, w)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userInfo, err := user.GetUserById(id)
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}
	err = utils.WriteJSON(w, http.StatusOK, userInfo)
	utils.ServerErrorHTTP(err, w)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData services.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}
	userCreated, err := user.CreateUser(userData)
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}
	err = utils.WriteJSON(w, http.StatusOK, userCreated)
	utils.ServerErrorHTTP(err, w)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userData services.User
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedUser, err := user.UpdateUser(id, userData)
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
	}
	err = utils.WriteJSON(w, http.StatusOK, updatedUser)
	utils.ServerErrorHTTP(err, w)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := user.DeleteUser(id)
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
	}
	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "Successfully deleted!"})
	utils.ServerErrorHTTP(err, w)
}
