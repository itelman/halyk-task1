package basic

import (
	"encoding/json"
	"net/http"
	"proxy-server/pkg/models"

	"github.com/gofrs/uuid"
)

func (app *Application) responseHandler(w http.ResponseWriter, r *http.Request) {
	var request models.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	httpReq, err := http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	for key, value := range request.Headers {
		httpReq.Header.Set(key, value)
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	defer httpResp.Body.Close()

	id, err := uuid.NewV4()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	response := models.Response{
		ID:      id.String(),
		Status:  httpResp.StatusCode,
		Headers: httpResp.Header,
		Length:  int(httpResp.ContentLength),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
