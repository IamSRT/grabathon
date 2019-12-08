package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"grabathon/api/request_response"
	"grabathon/models"
	"grabathon/service"
	"grabathon/util"
	"net/http"
	"strconv"
)

func CreateVouchHandler(w http.ResponseWriter, r *http.Request) {
	vouch := request_response.Vouch{}
	decodeErr := render.DecodeJSON(r.Body, &vouch)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	vouchResponse := service.CreateVouch(vouch)
	util.Send(w, r, "", vouchResponse)
}

func CreateVouchesHandler(w http.ResponseWriter, r *http.Request) {
	var vouch []request_response.Vouch
	decodeErr := render.DecodeJSON(r.Body, &vouch)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	vouchResponse := service.CreateVouches(vouch)
	util.Send(w, r, "", vouchResponse)
}


func GetVouchHandler(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "id")
	id, err := strconv.Atoi(t)
	if err != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}
	vouch := service.GetVouch(id)
	util.Send(w, r, "", vouch)
}

func GetAllVouchesForVoucheeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	vouch := service.GetAllVouchesForVouchee(id)
	util.Send(w, r, "", vouch)
}

func GetAllVouchesForVoucherHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	vouch := service.GetAllVouchesForVoucher(id)
	util.Send(w, r, "", vouch)
}

func DeleteVouchHandler(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "id")
	id, err := strconv.Atoi(t)
	if err != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}
	vouch := service.DeleteVouch(id)
	util.Send(w, r, "", vouch)
}

func UpdateVouchHandler(w http.ResponseWriter, r *http.Request) {
	vouch := request_response.Vouch{}
	decodeErr := render.DecodeJSON(r.Body, &vouch)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	vouchResponse, _ := service.UpdateVouch(vouch.Id, vouch.Status)
	util.Send(w, r, "", vouchResponse)
}

func AcceptVouch(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "id")
	id, err := strconv.Atoi(t)
	if err != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	vouchResponse, _ := service.UpdateVouch(id, models.Accepted)
	util.Send(w, r, "", vouchResponse)
}

func RejectVouch(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "id")
	id, err := strconv.Atoi(t)
	if err != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	vouchResponse, _ := service.UpdateVouch(id, models.Rejected)
	util.Send(w, r, "", vouchResponse)
}