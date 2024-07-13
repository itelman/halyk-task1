package main

import (
	"proxy-server/internal/app"
)

// @title HTTP Proxy Server API
// @version 1.0
// @description This is a server to proxy HTTP requests.
// @host localhost:8080
// @BasePath /
func main() {
	app.Run()
}
