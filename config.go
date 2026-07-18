package main

import (
	"github.com/dev-karani/FinTrack/internal/database"
	"github.com/dev-karani/fintrack/internal/database"
)

type api struct {
	db        *database.Queries
	jwtSecret string
	platform  string
}
