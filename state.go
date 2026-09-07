package main

import (
	"github.com/someshubham/gator/internal/config"
	"github.com/someshubham/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
