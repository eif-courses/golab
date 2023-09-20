package db

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	DB *sql.DB
}

var dbConnection = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDb(d)
	if err != nil {
		return nil, err
	}
	dbConnection.DB = d
	return dbConnection, err
}

func testDb(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	fmt.Println("Database Successfully Connected")
	return nil
}
