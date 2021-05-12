package main

import (
	"encoding/json"
	"net/http"
)

// HomeHandler serves the static file SPA
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("foo: bar")
}

// DetailHandler serves the
func DetailHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query()["url"][0]
	w.Header().Set("Content-Type", "application/json")
	body, err := QueryURL(u)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	response, _ := AssembleResponse(u, body)
	json.NewEncoder(w).Encode(response)
}
