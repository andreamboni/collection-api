package config

import "database/sql"

func InitializeMySQL() (*sql.DB, error) {
	logger := GetLogger("mysql")

	var err error
	db, err := sql.Open("mysql", MySQLConfig().FormatDSN())

	if err != nil {
		logger.Errorf("mysql opening error: %v", err)
		return nil, err
	}

	logger.Infof("mysql connected")
	return db, nil
}
