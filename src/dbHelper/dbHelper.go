package dbHelper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connect() {
	dbw := DbWorker{
		Dsn: "user:password@root(127.0.0.1:3306)/blog",
	}
	db, err := sql.Open("mysql",
		dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
}

