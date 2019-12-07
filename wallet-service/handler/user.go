package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"grabathon/api/request_response"
	"grabathon/service"
	"grabathon/util"
	"net/http"
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
	user := service.GetUser(userId)
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

func EnableAutoPay(w http.ResponseWriter, r *http.Request) {
	var vouches request_response.Vouches
	decodeErr := render.DecodeJSON(r.Body, &vouches)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	userResponse := service.EnableAutoPay(vouches.Vouches)
	util.Send(w, r, "", userResponse)
}

func Vouch(w http.ResponseWriter, r *http.Request) {
	var vouches request_response.Vouches
	decodeErr := render.DecodeJSON(r.Body, &vouches)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	vouchesResponse := service.Vouch(vouches.Vouches)
	util.Send(w, r, "", vouchesResponse)
}

func IsAutoPayEnabled(w http.ResponseWriter, r *http.Request) {
	vouch := request_response.Vouch{}
	decodeErr := render.DecodeJSON(r.Body, &vouch)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	vouchResponse := service.IsAutoEnabled(vouch.VoucheeId, vouch.VoucherId)
	util.Send(w, r, "", vouchResponse)
}

func IsVouchValid(w http.ResponseWriter, r *http.Request) {
	vouch := request_response.Vouch{}
	decodeErr := render.DecodeJSON(r.Body, &vouch)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	vouchResponse := service.IsVouchValid(vouch.VoucheeId, vouch.VoucherId, vouch.Amount)
	util.Send(w, r, "", vouchResponse)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	util.Send(w, r, "", users)
}