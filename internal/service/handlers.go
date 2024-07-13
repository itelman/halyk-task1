package service

import (
	"encoding/json"
	"net/http"
	"proxy-server/pkg/models"
)

// responseHandler godoc
//
//	@Summary		Proxy HTTP request
//	@Description	Proxies an HTTP request to a specified URL.
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.Request	true	"Request payload"
//	@Success		200		{object}	models.Response
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/ [post]
func (app *Application) responseHandler(w http.ResponseWriter, r *http.Request) {
	var request models.Request
	var response models.Response

	if err := app.readJSON(w, r, &request); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := response.Set(request); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.Storage.Store(response.ID, envelope{"request": request, "response": response})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// healthCheckHandler godoc
//
//	@Summary		Health check
//	@Description	This endpoint checks the health of the server.
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"OK"
//	@Router			/health [get]
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
