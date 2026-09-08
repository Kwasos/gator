package main

import (
	"github.com/Kwasos/gator/internal/config"
	"github.com/Kwasos/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
