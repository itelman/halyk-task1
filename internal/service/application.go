package service

import (
	"proxy-server/internal/jsonlog"
	"sync"
)

type Application struct {
	Config Config
	Logger *jsonlog.Logger
	WG     sync.WaitGroup
}

const version = "0.0.1"
