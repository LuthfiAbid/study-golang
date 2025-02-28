package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(host:3306)/test-golang")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
