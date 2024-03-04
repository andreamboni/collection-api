package handler

import (
	"database/sql"

	"collection.com/config"
)

var (
	db     *sql.DB
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMySQL()
}
