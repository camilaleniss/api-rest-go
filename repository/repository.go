package repository

import (
	"github.com/camilaleniss/api-rest-go/connection"
)

type DomainRepo interface {
	CreateDomain(host string, ssl_grade string, ssl_previous_grade string)
	SearchDomain(hostquery string) (connection.DomainBD, error)
	SearchDomains() []connection.DomainBD
	UpdateDomain(host string, new_ssl_grade string, new_previous_grade string) error
	UpdateDomainVisit(host string) error
}
