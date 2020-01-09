package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/camilaleniss/api-rest-go/connection"
	dRepo "github.com/camilaleniss/api-rest-go/repository"
	_ "github.com/lib/pq"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewSQLDomainRepo(Conn *sql.DB) dRepo.DomainRepo {
	return &sqlDomainRepo{
		Conn: Conn,
	}
}

type sqlDomainRepo struct {
	Conn *sql.DB
}

func (s *sqlDomainRepo) CreateDomain(host string, ssl_grade string, ssl_previous_grade string) {
	values := "'" + host + "','" + ssl_grade + "','" + ssl_previous_grade + "', "
	if _, err := s.Conn.Exec(
		"INSERT INTO domain (host, ssl_grade, ssl_previous_grade, last_search) VALUES (" + values + " NOW())"); err != nil {
		log.Fatal(err)
	}
}

func (s *sqlDomainRepo) SearchDomain(hostquery string) (connection.DomainBD, error) {

	query := "SELECT host, ssl_grade, ssl_previous_grade, last_search FROM domain WHERE host='" + hostquery + "';"
	rows, err := s.Conn.Query(query)
	if err != nil {
		return connection.DomainBD{}, err
	}

	defer rows.Close()
	var host, ssl_grade, ssl_previous_grade string
	var last_search time.Time

	for rows.Next() {
		if err := rows.Scan(&host, &ssl_grade, &ssl_previous_grade, &last_search); err != nil {
			log.Fatal(err)
		}

		return connection.DomainBD{host, ssl_grade, ssl_previous_grade, last_search}, err
	}

	return connection.DomainBD{}, err
}

func (s *sqlDomainRepo) SearchDomains() []connection.DomainBD {
	query := "SELECT host FROM domain ORDER BY last_search DESC;"
	rows, err := s.Conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var hosts []connection.DomainBD
	var host string

	for rows.Next() {
		if err := rows.Scan(&host); err != nil {
			log.Fatal(err)
		}

		domain := connection.DomainBD{}
		domain.Host = host
		hosts = append(hosts, domain)
	}

	return hosts
}

func (s *sqlDomainRepo) UpdateDomain(host string, new_ssl_grade string, new_previous_grade string) error {
	if _, err := s.Conn.Exec(
		"UPDATE domain SET ssl_grade =  $1, ssl_previous_grade = $2 , last_search = NOW() WHERE host = $3", new_ssl_grade,
		new_previous_grade, host); err != nil {
		return err
	}
	return nil
}
