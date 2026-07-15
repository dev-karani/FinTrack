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

	//pass arg to routes func
	registerRoutes(mux)
	// init server struct and values
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//start server
	log.Fatal(server.ListenAndServe())
}
