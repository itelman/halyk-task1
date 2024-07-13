package service

import (
	"net/http"

	_ "proxy-server/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			HTTP Proxy Server API
// @version		1.0
// @description	This is a server to proxy HTTP requests.
// @host			localhost:8080
// @BasePath		/
func (app *Application) routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/health", app.healthCheckHandler).Methods("GET")
	router.HandleFunc("/", app.responseHandler).Methods("POST")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)

	return router
}
