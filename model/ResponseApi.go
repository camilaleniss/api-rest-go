package model

import (
	"log"

	"github.com/likexian/whois-go"
	_ "github.com/likexian/whois-go"
	//"github.com/domainr/whois"
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

func WhoisServerAttributes(server ServerApi) string {

	ip := server.IpAddress
	result, err := whois.Whois(ip)
	if err != nil {
		log.Fatal(err)
	}
	return result

	/*
		request, err := whois.NewRequest(ip)
		if err != nil {
			log.Fatal(err)
		}
		response, err := whois.DefaultClient.Fetch(request)
		if err != nil {
			log.Fatal(err)
		}
		return response.String()
	*/
}
