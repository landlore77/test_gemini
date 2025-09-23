package utils

import (
	"database/sql"
	"fmt"
	"test2/utils/config"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	cfg := config.AppConfig.MYSQL
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/test_admin", cfg.ID, cfg.PASS)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	return db, nil
}
