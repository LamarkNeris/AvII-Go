package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Sql() *sql.DB {
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/AV II GO")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return database
}
