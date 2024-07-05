package main

import (
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	/*s, err := app.snippets.Get(id)
	if err != nil {
		app.serverError(w, err)
		return
	}*/

	var data []byte

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
