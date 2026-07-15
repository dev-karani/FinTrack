package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("fintrack journey begins!")

	//	initialise server handler
	mux := http.NewServeMux()

	// register routes
	registerRoutes(mux)

	// init server struct
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// run server
	log.Fatal(server.ListenAndServe())
}
