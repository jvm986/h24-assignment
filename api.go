package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// URLHandler serves the endpoint for requesting url info
func URLHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query()["url"][0]
	re := regexp.MustCompile("^(http://|https://)")
	if !re.Match([]byte(u)) {
		u = "https://" + u
	}
	fmt.Printf("Processing %s\n", u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	body, err := QueryURL(u)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	response, _ := AssembleResponse(u, body)
	json.NewEncoder(w).Encode(response)
}
