package statistics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)


//func TestGetStatsPerDay(t *testing.T) {
//	AddCustomRules()
//	var URL *url.URL
//	URL, _ = url.Parse("http://www.example.com")
//	params := url.Values{}
//	params.Add("limit", "1")
//	params.Add("date1", "2014-06-20")
//	params.Add("date2", "2014-06-20")
//	params.Add("action", "like")
//	URL.RawQuery = params.Encode()
//	r, _ := http.NewRequest("GET", URL.String(), nil)
//	e := validateGetStats(r)
//	if len(e) != 0 {
//		fmt.Println(e)
//		t.Error("incorrect behavior of validateGetStats")
//	}
//	params = url.Values{}
//	params.Add("action", "like1")
//	URL.RawQuery = params.Encode()
//	r, _ = http.NewRequest("GET", URL.String(), nil)
//	e = validateGetStats(r)
//	if len(e) == 0 {
//		fmt.Println(e)
//		t.Error("incorrect behavior of validateGetStats")
//	}
//}
//
//func TestCreateUser(t *testing.T) {
//	AddCustomRules()
//	var URL *url.URL
//	URL, _ = url.Parse("http://www.example.com")
//	user := &User{Id: 1,
//		Sex:"M", Age:14}
//	body, _ := json.Marshal(user)
//	r, _ := http.NewRequest("GET", URL.String(), bytes.NewReader(body))
//	_, e := validateCreateUser(r)
//	if len(e) != 0 {
//		fmt.Println(e)
//		t.Error("incorrect behavior of validateCreateUser")
//	}
//	user = &User{Id: 1, Age:14}
//	body, _ = json.Marshal(user)
//	r, _ = http.NewRequest("GET", URL.String(), bytes.NewReader(body))
//	_, e = validateCreateUser(r)
//	if len(e) == 0 {
//		fmt.Println(e)
//		t.Error("incorrect behavior of validateCreateUser")
//	}
//}

func TestCreateStat(t *testing.T) {
	AddCustomRules()
	var URL *url.URL
	URL, _ = url.Parse("http://www.example.com")
	user := &Stat{User: 1,
		Action:"like", Ts:"2018-11-30T18:12:34"}
	body, _ := json.Marshal(user)
	r, _ := http.NewRequest("GET", URL.String(), bytes.NewReader(body))
	_, e := validateCreateStat(r)
	if len(e) != 0 {
		fmt.Println(e)
		t.Error("incorrect behavior of validateCreateStat")
	}
	user = &Stat{User: 1, Ts:"2018-11-30T18:12:34"}
	body, _ = json.Marshal(user)
	r, _ = http.NewRequest("GET", URL.String(), bytes.NewReader(body))
	_, e = validateCreateStat(r)
	if len(e) == 0 {
		fmt.Println(e)
		t.Error("incorrect behavior of validateCreateUser")
	}
}