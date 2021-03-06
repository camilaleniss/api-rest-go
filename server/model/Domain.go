package model

type Domain struct {
	//Host             string   `json:"host"`
	Servers          []Server `json:"servers"`
	ServersChanged   bool     `json:"servers_changed"`
	SslGrade         string   `json:"ssl_grade"`
	PreviousSslGrade string   `json:"previous_ssl_grade"`
	Logo             string   `json:"logo"`
	Title            string   `json:"title"`
	IsDown           bool     `json:"is_down"`
}

type Server struct {
	Address  string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
}

type DomainResponse struct {
	Domains []HostResponse `json:"items"`
}

type HostResponse struct {
	Host string `json:"host"`
}
