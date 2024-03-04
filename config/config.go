package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB

	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeMySQL()

	if err != nil {
		return fmt.Errorf("error initializing mysql: %v", err)
	}

	return nil
}

func GetMySQL() *sql.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

func MySQLConfig() *mysql.Config {
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "collection",
		ParseTime: true,
	}

	return &cfg
}
