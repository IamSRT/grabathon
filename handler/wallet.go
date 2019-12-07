package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"grabathon/api/request_response"
	"grabathon/service"
	"grabathon/util"
	"net/http"
)

func CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	wallet := request_response.Wallet{}
	decodeErr := render.DecodeJSON(r.Body, &wallet)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	walletResponse := service.CreateWallet(wallet)
	util.Send(w, r, "", walletResponse)
}


func GetWalletHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	wallet := service.GetWallet(userId)
	util.Send(w, r, "", wallet)
}

func UpdateWalletHandler(w http.ResponseWriter, r *http.Request) {
	wallet := request_response.Wallet{}
	decodeErr := render.DecodeJSON(r.Body, &wallet)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	walletResponse := service.UpdateWallet(wallet)
	util.Send(w, r, "", walletResponse)
}