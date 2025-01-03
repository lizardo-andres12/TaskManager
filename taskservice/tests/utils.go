package tests

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func loadDB() (*sql.DB, string) {
	err := godotenv.Load()
	if err != nil {
		return nil, "Error loading .env file"
	}

	var db *sql.DB
	cfg := mysql.Config{
		User:   os.Getenv("TEST_DB_USER"),
		Passwd: os.Getenv("TEST_DB_PASSWD"),
		Net:    os.Getenv("TEST_DB_NET"),
		Addr:   os.Getenv("TEST_DB_ADDR"),
		DBName: os.Getenv("TEST_DB_NAME"),
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, "Invalid login to test database"
	}
	return db, ""
}
