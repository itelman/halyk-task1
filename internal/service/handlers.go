package service

import (
	"net/http"
)

func (app *Application) responseHandler(w http.ResponseWriter, r *http.Request) {
	var data []byte

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.Config.Env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
