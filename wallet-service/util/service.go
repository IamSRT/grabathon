package util

import (
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

/**
 * Created by Sai Ravi Teja K on 22, Jul 2019
 */

// Helper function to print all the routes of application
func PrintAllRoutes(router *chi.Mux) {
	walkAllRoutesFunc := func(method string, route string,
		_ http.Handler, _ ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)         // walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkAllRoutesFunc); err != nil {
		log.Warnf("Logging err: %s\n", err.Error()) // panic if there is an error
	}
}