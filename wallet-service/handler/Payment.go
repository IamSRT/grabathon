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

func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	payment := request_response.Payment{}
	decodeErr := render.DecodeJSON(r.Body, &payment)
	if decodeErr != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	paymentResponse := service.CreatePayment(payment)
	util.Send(w, r, "", paymentResponse)
}


func GetPaymentHandler(w http.ResponseWriter, r *http.Request) {
	pId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(pId)
	if err != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	payment := service.GetPayment(id)
	util.Send(w, r, "", payment)
}

func GetPendingPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	uId := chi.URLParam(r, "id")
	completeWallet := service.GetPendingPayments(uId)
	util.Send(w, r, "", completeWallet)
}

func DeletePaymentHandler(w http.ResponseWriter, r *http.Request) {
	pId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(pId)
	if err != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}

	err = service.DeletePayment(id)
	if err != nil {
		util.SendInternalServerError(w, r, "Failed to parse input request")
		return
	}
}