package service

import (
	"net/http"

	_ "proxy-server/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			HTTP Proxy Server API
//	@version		1.0
//	@description	This is a server to proxy HTTP requests.
//	@host			localhost:8080
//	@BasePath		/
func (app *Application) routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/health", app.healthCheckHandler)
	router.HandleFunc("/", app.responseHandler)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
