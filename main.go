package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"grabathon/config"
	"grabathon/logger"
	"grabathon/routes"
	"grabathon/util"
	"net/http"
)

/**
 * Created by Sai Ravi Teja K on 22, Jul 2019
 */

func main() {

	router := getRouter()
	util.PrintAllRoutes(router)

	url := ":" + config.ServerPort
	server := &http.Server{Addr: url, Handler: router}
	err := server.ListenAndServe()

	if err != nil {
		log.Warnf("Logging err: %s\n", err.Error())
	}
}


func getRouter() *chi.Mux {
	router := chi.NewRouter()

	requestLogger := log.New()
	requestLogger.Formatter = &log.JSONFormatter{}

	router.Use(
		render.SetContentType(render.ContentTypeJSON), 			// Set content-Type headers as application/json
		logger.StructuredRequestLogger(requestLogger),			// Log API request calls
		middleware.DefaultCompress,                    			// Compress results, mostly g-zipping assets and json
		middleware.RedirectSlashes,                    			// Redirect slashes to no slash URL versions
		middleware.Recoverer,                          			// Recover from panics without crashing server
	)

	routes.ConfigureRoutes(router)
	return router
}