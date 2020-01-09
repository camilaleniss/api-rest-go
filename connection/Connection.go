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

/*
func CreateDomain(db *sql.DB, host string, ssl_grade string, ssl_previous_grade string) {
	values := "'" + host + "','" + ssl_grade + "','" + ssl_previous_grade + "', "
	if _, err := db.Exec(
		"INSERT INTO domain (host, ssl_grade, ssl_previous_grade, last_search) VALUES (" + values + " NOW())"); err != nil {
		log.Fatal(err)
	}
}

func SearchDomain(db *sql.DB, hostquery string) DomainBD {

	query := "SELECT host, ssl_grade, ssl_previous_grade, last_search FROM domain WHERE host='" + hostquery + "';"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var host, ssl_grade, ssl_previous_grade string
	var last_search time.Time

	for rows.Next() {
		if err := rows.Scan(&host, &ssl_grade, &ssl_previous_grade, &last_search); err != nil {
			log.Fatal(err)
		}

		return DomainBD{host, ssl_grade, ssl_previous_grade, last_search}
	}

	return DomainBD{}
}

func SearchDomains(db *sql.DB) []DomainBD {
	query := "SELECT host FROM domain ORDER BY last_search DESC;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var hosts []DomainBD
	var host string

	for rows.Next() {
		if err := rows.Scan(&host); err != nil {
			log.Fatal(err)
		}

		domain := DomainBD{}
		domain.Host = host
		hosts = append(hosts, domain)
	}

	return hosts
}

func UpdateDomain(db *sql.DB, host string, new_ssl_grade string, new_previous_grade string) error {
	if _, err := db.Exec(
		"UPDATE domain SET ssl_grade =  $1, ssl_previous_grade = $2 , last_search = NOW() WHERE host = $3", new_ssl_grade,
		new_previous_grade, host); err != nil {
		return err
	}
	return nil
}
*/
