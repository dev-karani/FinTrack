package main

import (
	"fmt"
	"net/http"
)

func registerRoutes(mux *http.ServeMux) {

	// root path
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<h1>Fintrack journey is live</h1>")
	})

	// Health check route
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "server is up and running")
	})

}
