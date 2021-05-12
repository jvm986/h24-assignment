package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/url", DetailHandler)

	http.ListenAndServe(":8000", nil)
}
