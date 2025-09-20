package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"test1/config"
)

func GetDB() (*sql.DB, error) {
	log.Printf("Attempting to connect to MySQL with User: %s", config.Cfg.MySQL.ID)
	cfg := mysql.Config{
		User:   config.Cfg.MySQL.ID,
		Passwd: config.Cfg.MySQL.PASS,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "test_admin",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, nil
}
