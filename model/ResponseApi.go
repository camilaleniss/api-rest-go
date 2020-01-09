package model

import (
	"log"
	"strings"

	"github.com/likexian/whois-go"
)

const SSL_DEFAULT = "-"
const LINE_OWNER = 41
const LINE_COUNTRY = 47

//https://github.com/ssllabs/research/wiki/SSL-Server-Rating-Guide
var ssl_grades = []string{"A", "B", "C", "D", "E", "F"}

type DomainApi struct {
	Host      string      `json:"host"`
	Endpoints []ServerApi `json:"endpoints"`
	Erros     []ErrorsApi `json:"errors"`
}

type ServerApi struct {
	IpAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
}

type ErrorsApi struct {
	Message string `json:"message"`
}

func WhoisServerAttributes(server ServerApi) (string, string) {

	ip := server.IpAddress
	result, err := whois.Whois(ip)
	if err != nil {
		log.Fatal(err)
	}

	owner, country := splitWhois(result)

	return owner, country

}

func splitWhois(response string) (string, string) {
	responses1 := (strings.Split(response, "\n"))

	lineOwner := responses1[LINE_OWNER]
	lineCountry := responses1[LINE_COUNTRY]

	owner := strings.Split(lineOwner, ":")[1]
	country := strings.Split(lineCountry, ":")[1]

	serverOwner := strings.Trim(owner, " ")
	serverCountry := strings.Trim(country, " ")

	return serverOwner, serverCountry
}

/*
The SSL grades of the servers goes from A to F where A is the biggest grade.
The SSL grade of a domain is the minor SSL grade of the servers
*/

func GenerateSSLGrade(servers []ServerApi) string {

	minor := servers[0].Grade

	for i := 1; i < len(servers); i++ {
		grade := strings.Split(servers[i].Grade, "")[0]
		if existsInSSL(grade) {
			if grade > minor {
				minor = grade
			}
		}
	}
	return minor
}

func existsInSSL(grade string) bool {
	for i := 0; i < len(ssl_grades); i++ {
		ssl := ssl_grades[i]
		if ssl == grade {
			return true
		}
	}
	return false
}
