package connection

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	Cock *sql.DB
}

var dbConn = &DB{}

type DomainBD struct {
	Host               string
	Ssl_grade          string
	Ssl_previous_grade string
	Last_search        time.Time
}

func GetConnection() (*DB, error) {
	dataBase, err := sql.Open("postgres", "postgresql://root@localhost:26257/domains?sslmode=disable")
	dbConn.Cock = dataBase

	/*
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	*/
	return dbConn, err
}
