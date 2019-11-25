package util

import (
	"github.com/go-chi/render"
	"grabathon/api/response"
	"net/http"
)

func SendInternalServerError(w http.ResponseWriter, r *http.Request, message string) {
	if len(message) == 0 {
		message = "Internal Server Error"
	}
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, response.ApiResponse{
		ApiMessage : message,
	})
}

func SendBadRequest(w http.ResponseWriter, r *http.Request, message string){
	if len(message) == 0 {
		message = "Bad Request"
	}
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, response.ApiResponse{
		ApiMessage : message,
	})
}

func SendForbidden(w http.ResponseWriter, r *http.Request, message string){
	if len(message) == 0 {
		message = "Success"
	}
	render.Status(r, http.StatusForbidden)
	render.JSON(w, r, response.ApiResponse{})
}

func SendNotFound(w http.ResponseWriter, r *http.Request){
	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, response.ApiResponse{})
}

func Send(w http.ResponseWriter, r *http.Request, message string, data interface{}){
	if len(message) == 0 {
		message = "Success"
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response.ApiResponse{
		ApiMessage : message,
		Data       : data,
	})
}

func SendCreated(w http.ResponseWriter, r *http.Request, message string, data interface{}){
	if len(message) == 0 {
		message = "Created"
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, response.ApiResponse{
		ApiMessage : message,
		Data       : data,
	})
}

