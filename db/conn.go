package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "user"
	password = "1234"
	dbname   = "mysql"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to " + dbname)

	return db, nil
}
