package database

import (
	"database/sql"
	"time"
)

func GetConn() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(host:3306)/test-golang")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}
