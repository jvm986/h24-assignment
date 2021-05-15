package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)
	http.HandleFunc("/url", URLHandler)
	fmt.Println("Serving on localhost port 8000... ")
	http.ListenAndServe(":8000", nil)
}
