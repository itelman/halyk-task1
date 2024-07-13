package service

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *Application) routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/health", app.healthCheckHandler)
	router.HandleFunc("/", app.responseHandler)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
