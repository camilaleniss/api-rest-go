package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	dataBase, err := sql.Open("postgres", "postgresql://root@localhost:26257/domains?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return dataBase
}
