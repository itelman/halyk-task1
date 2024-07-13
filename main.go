package main

import (
	"os"
	"proxy-server/internal/app"
)

var port = os.Getenv("PORT")

func main() {
	app.Run()
}
