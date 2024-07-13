package app

import (
	"flag"
	"os"
	"proxy-server/internal/service"
	"proxy-server/internal/service/jsonlog"
)

func Run() {
	var cfg service.Config

	flag.IntVar(&cfg.Port, "port", 8080, "port for api")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &service.Application{
		Config: cfg,
		Logger: logger,
	}

	err := app.Serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
