package tests

import (
	"fmt"
	"testing"
	"time"

	handler "github.com/camilaleniss/api-rest-go/handler/http"
	"github.com/camilaleniss/api-rest-go/model"
)

func SetUp9() []model.ServerApi {
	return []model.ServerApi{}
}

func SetUp10() []model.ServerApi {
	servers := []model.ServerApi{}
	servers = append(servers, model.ServerApi{"", "A"})
	return servers
}

func TestCompareOneHourBefore(t *testing.T) {

	var servers []model.ServerApi
	//Test1
	last_search := time.Now()
	servers = SetUp9()

	ssl, prev := handler.CompareOneHourBefore(servers, "A", last_search)
	fmt.Println(ssl + "    " + prev)

	if ssl != prev {
		t.Errorf("It must be equal, cause it haven´t passed one hour yet")
	}

	//Test2
	servers = SetUp10()
	ssl, prev = handler.CompareOneHourBefore(servers, "A", last_search)
	fmt.Println(ssl + "    " + prev)

	if ssl != prev {
		t.Errorf("It must be equal, cause it haven´t passed one hour yet")
	}

	//Test3
	last_search = time.Now()
	last_search = last_search.Add(-2 * time.Hour)
	servers = SetUp9()

	ssl, prev = handler.CompareOneHourBefore(servers, "B", last_search)

	fmt.Println(ssl + "    " + prev)

	if ssl != "-" {
		t.Errorf("It must be - cause is the updated")
	}

	if prev != "B" {
		t.Errorf("It must be B cause is the updated")
	}

	//Test4
	servers = SetUp10()
	ssl, prev = handler.CompareOneHourBefore(servers, "B", last_search)

	fmt.Println(ssl + "    " + prev)

	if ssl != "A" {
		t.Errorf("It must be A cause is the updated")
	}

	if prev != "B" {
		t.Errorf("It must be B cause is the updated")
	}

}

func TestGetHTML(t *testing.T) {
	url := "https://www.truora.com/"
	title2, icon := handler.GetHTMLInfo(url)
	fmt.Println(title2)
	fmt.Println(icon)
}
