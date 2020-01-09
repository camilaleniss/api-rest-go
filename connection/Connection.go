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

func CreateDomain(host string, ssl_grade string, ssl_previous_grade string) {
	db := GetConnection()

	values := "'" + host + "','" + ssl_grade + "','" + ssl_previous_grade + "', "
	/*
		if _, err := db.Exec(
			`INSERT INTO domain (host, ssl_grade, ssl_previous_grade, last_search)
				VALUES (?,?,?,?) ` + values); err != nil {
			log.Fatal(err)

		}
	*/

	if _, err := db.Exec(
		"INSERT INTO domain (host, ssl_grade, ssl_previous_grade, last_search) VALUES (" + values + " NOW())"); err != nil {
		log.Fatal(err)
	}
}
