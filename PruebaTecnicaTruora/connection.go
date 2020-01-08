package main

import (
	"database/sql"
	"log"
)

func GetConnection() *sql.DB {
	dataBase, err := sql.Open("postgres", "postgresql://root@localhost:26257/domains?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return dataBase
}
