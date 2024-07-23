package config

import (
	"net/http"
	"proxy-server/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "proxy-server/docs"
)

// @title			HTTP Proxy Server API
// @version		1.0
// @description	This is a server to proxy HTTP requests.
// @host			localhost:8080
// @BasePath		/
func Routes() http.Handler {
	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our app receives.
	standardMiddleware := alice.New(recoverPanic, logRequest, secureHeaders)

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.ResponseHandler).Methods("POST")

	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowedResponse)

	return standardMiddleware.Then(router)
}
