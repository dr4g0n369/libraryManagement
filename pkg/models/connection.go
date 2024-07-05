package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func makeConnectionURL() string {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbIPAddress := os.Getenv("DB_IPADDR")
	// dbDrvier := os.Getenv("DB_DRIVER")
	return fmt.Sprintf("%v:%v@(%v:%v)/%v?parseTime=true", dbUser, dbPass, dbIPAddress, dbPort, dbName)
}

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", makeConnectionURL())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
