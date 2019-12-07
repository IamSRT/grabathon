package handler

import (
	"grabathon/service"
	"grabathon/util"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	serviceHealth := service.GetServiceHealth()
	util.Send(w, r, "Health check passing", serviceHealth)
}
