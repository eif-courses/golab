package controllers

import (
	"encoding/json"
	"github.com/eif-courses/golab/services"
	"github.com/eif-courses/golab/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

// User godoc
var user services.User

type User struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	all, err := user.GetAllUsers()
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}
	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"users": all})
	utils.ServerErrorHTTP(err, w)
}

// GetUserById godoc
// @Summary Get user by id
// @Description Get user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Router /users/user/{id} [get]
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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "User"
// @Success 200 {object} User
// @Router /users/user [post]
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

// UpdateUser godoc
// @Summary Update user by id
// @Description Update user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body User true "User"
// @Success 200 {object} User
// @Router /users/user/{id} [put]
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

// DeleteUser godoc
// @Summary Delete user by id
// @Description Delete user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Router /users/user/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := user.DeleteUser(id)
	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
	}
	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "Successfully deleted!"})
	utils.ServerErrorHTTP(err, w)
}
