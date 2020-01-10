package tests

import (
	"testing"

	"github.com/camilaleniss/api-rest-go/model"
)

func SetUp1() []model.ServerApi {
	return []model.ServerApi{}
}

func SetUp2() []model.ServerApi {
	servers := []model.ServerApi{}
	servers = append(servers, model.ServerApi{"", "A"})
	return servers
}

func SetUp3() []model.ServerApi {
	servers := []model.ServerApi{}
	servers = append(servers, model.ServerApi{"", "A"})
	servers = append(servers, model.ServerApi{"", "B"})
	servers = append(servers, model.ServerApi{"", "C"})
	return servers
}

func SetUp4() []model.ServerApi {
	servers := []model.ServerApi{}
	servers = append(servers, model.ServerApi{"", "C"})
	servers = append(servers, model.ServerApi{"", "B"})
	servers = append(servers, model.ServerApi{"", "A"})
	return servers
}

func SetUp5() []model.ServerApi {
	servers := []model.ServerApi{}
	servers = append(servers, model.ServerApi{"", "B"})
	servers = append(servers, model.ServerApi{"", "C"})
	servers = append(servers, model.ServerApi{"", "C"})
	return servers
}

func SetUp6() []model.ServerApi {
	servers := []model.ServerApi{}
	servers = append(servers, model.ServerApi{"", "-"})
	servers = append(servers, model.ServerApi{"", "-"})
	servers = append(servers, model.ServerApi{"", "-"})
	return servers
}

func SetUp7() []model.ServerApi {
	servers := []model.ServerApi{}
	servers = append(servers, model.ServerApi{"", "A"})
	servers = append(servers, model.ServerApi{"", "-"})
	servers = append(servers, model.ServerApi{"", "-"})
	return servers
}

func SetUp8() []model.ServerApi {
	servers := []model.ServerApi{}
	servers = append(servers, model.ServerApi{"", "B"})
	servers = append(servers, model.ServerApi{"", "B"})
	servers = append(servers, model.ServerApi{"", "-"})
	return servers
}

func TestGenerateSslGrade(t *testing.T) {
	var servers []model.ServerApi
	var ssl_grade string

	//Test1
	servers = SetUp1()
	ssl_grade = model.GenerateSSLGrade(servers)

	if ssl_grade != model.SSL_DEFAULT {
		t.Errorf("It must be - , is an empty array")
	}

	//Test2
	servers = SetUp2()
	ssl_grade = model.GenerateSSLGrade(servers)

	if ssl_grade != "A" {
		t.Errorf("It must be A")
	}

	//Test3
	servers = SetUp3()
	ssl_grade = model.GenerateSSLGrade(servers)

	if ssl_grade != "C" {
		t.Errorf("It must be C")
	}

	//Test4
	servers = SetUp4()
	ssl_grade = model.GenerateSSLGrade(servers)

	if ssl_grade != "C" {
		t.Errorf("It must be C")
	}

	//Test5
	servers = SetUp5()
	ssl_grade = model.GenerateSSLGrade(servers)

	if ssl_grade != "C" {
		t.Errorf("It must be C")
	}

	//Test6
	servers = SetUp6()
	ssl_grade = model.GenerateSSLGrade(servers)

	if ssl_grade != model.SSL_DEFAULT {
		t.Errorf("It must be - all the servers has that grade")
	}

	//Test7
	servers = SetUp7()
	ssl_grade = model.GenerateSSLGrade(servers)

	if ssl_grade != "A" {
		t.Errorf("It must be A")
	}

	//Test8
	servers = SetUp8()
	ssl_grade = model.GenerateSSLGrade(servers)

	if ssl_grade != "B" {
		t.Errorf("It must be B")
	}
}

func TestWhoIs(t *testing.T) {
	ip := "2606:4700:0:0:0:0:6810:787f"
	owner, country := model.WhoisServerAttributes(model.ServerApi{IpAddress: ip})

	if owner != "" || country != "" {
		t.Errorf("It has no way to come this far")
	}

}
