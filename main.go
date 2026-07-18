package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dev-karani/FinTrack/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("fintrack journey begins!")

	if err := godotenv.Load(); err != nil {
		log.Printf("warning: could not load .env: %v", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("APP_ENV")

	if jwtSecret == "" {
		log.Fatal("missing jwt secret")
	}
	if platform == "" {
		log.Fatal("platform must be set")
	}

	// creat db pool
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	cfg := &apiConfig{
		db:        database.New(db),
		platform:  platform,
		jwtSecret: jwtSecret,
	}

	//	initialise server handler
	mux := http.NewServeMux()

	// register routes
	registerRoutes(mux, cfg)

	// init server struct
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// run server
	log.Fatal(server.ListenAndServe())
}
