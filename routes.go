package main

import (
	"fmt"
	"net/http"

	transactions "github.com/dev-karani/FinTrack/transactions"
	users "github.com/dev-karani/FinTrack/users"
)

const apiV1 = "/api/v1"

func registerRoutes(
	mux *http.ServeMux,
	usersHandler *users.Handler,
	transactionsHandler *transactions.Handler,
) {
	// Root
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<h1>FinTrack journey is live </h1>")
	})

	// Health
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "server is up and running")
	})

	// User routes
	mux.HandleFunc("POST "+apiV1+"/users", usersHandler.CreateUser)
	mux.HandleFunc("POST "+apiV1+"/login", usersHandler.Login)
	mux.HandleFunc("POST "+apiV1+"/refresh", usersHandler.Refresh)
	mux.HandleFunc("POST "+apiV1+"/revoke", usersHandler.Revoke)

	mux.HandleFunc("PUT "+apiV1+"/users", usersHandler.UpdateUser)
	mux.HandleFunc("DELETE "+apiV1+"/users", usersHandler.DeleteUser)

	// Transaction routes

	mux.HandleFunc("POST "+apiV1+"/transactions", transactionsHandler.CreateTransaction)
	mux.HandleFunc("GET "+apiV1+"/transactions", transactionsHandler.GetAllUserTransactions)

	mux.HandleFunc("GET "+apiV1+"/transactions/{transactionID}", transactionsHandler.GetTransactionByID)
	mux.HandleFunc("PUT "+apiV1+"/transactions/{transactionID}", transactionsHandler.UpdateTransaction)
	mux.HandleFunc("DELETE "+apiV1+"/transactions/{transactionID}", transactionsHandler.DeleteTransactionByID)
}
