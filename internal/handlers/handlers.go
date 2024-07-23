package handlers

import (
	"encoding/json"
	"net/http"
	"proxy-server/internal/service/helpers"
	"proxy-server/pkg/models"
	"sync"
)

var store sync.Map

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
func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	var request models.Request

	if err := helpers.ReadJSON(w, r, &request); err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	response, err := models.NewResponse(request)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	store.Store(response.ID, map[string]interface{}{"request": request, "response": response})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
