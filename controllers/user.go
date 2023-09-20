package controllers

import (
	"encoding/json"
	"github.com/eif-courses/golab/helpers"
	"github.com/eif-courses/golab/services"
	"net/http"
)

// GET/users

var user services.User

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	all, err := user.GetAllUsers()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"users": all})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData services.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	userCreated, err := user.CreateUser(userData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, userCreated)
}
