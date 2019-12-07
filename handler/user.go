package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"grabathon/api/request_response"
	"grabathon/service"
	"grabathon/util"
	"net/http"
	"strconv"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := request_response.User{}
	decodeErr := render.DecodeJSON(r.Body, &user)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	userResponse := service.CreateUser(user)
	util.Send(w, r, "", userResponse)
}


func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}
	user := service.GetUser(id)
	util.Send(w, r, "", user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := request_response.User{}
	decodeErr := render.DecodeJSON(r.Body, &user)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	userResponse := service.UpdateUser(user)
	util.Send(w, r, "", userResponse)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	util.Send(w, r, "", users)
}