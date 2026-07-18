package main

import (
	"github.com/dev-karani/FinTrack/internal/database"
)

type api struct {
	db        *database.Queries
	jwtSecret string
	platform  string
}
