package model

import (
	"log"
	"strings"

	"github.com/likexian/whois-go"
)

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

	lineOwner := responses1[41]
	lineCountry := responses1[47]

	owner := strings.Split(lineOwner, ":")[1]
	country := strings.Split(lineCountry, ":")[1]

	serverOwner := strings.Trim(owner, " ")
	serverCountry := strings.Trim(country, " ")

	return serverOwner, serverCountry
}

func GenerateSSLGrade(servers []ServerApi) string {
	return ""
}
