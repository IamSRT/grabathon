package handler

import (
	"github.com/go-chi/render"
	"grabathon/api/request"
	"grabathon/service"
	"grabathon/util"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := request.User{}
	decodeErr := render.DecodeJSON(r.Body, &user)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	userResponse := service.CreateUser(user)
	util.Send(w, r, "", userResponse)
}


func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := service.GetUser()
	util.Send(w, r, "", user)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	util.Send(w, r, "", users)
}