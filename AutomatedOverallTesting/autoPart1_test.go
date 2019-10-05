package AutomatedOverallTesting

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"

	"TuriteaWebResources/server/actions"
)

func init() {
	go func() {
		config := &actions.Config{true, true}
		actions.Start(config)
		log.Println("server start")
		for {
			err := http.ListenAndServe("localhost:80", nil)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	client.Client = &http.Client{}
	var err error
	u, err = url.Parse("http://localhost")
	if err != nil {
		panic(err)
	}
	client.Jar, err = cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	cookieInit = []*http.Cookie{
		{Name:"super", MaxAge:0, Path:"/"},
		{Name:"key", MaxAge:0, Path:"/"},
		{Name:"lastTime", MaxAge:0, Path:"/"},
	}
}
var u *url.URL
var client = &myClient{}
var cookieInit []*http.Cookie

type myClient struct {
	*http.Client
}

func TestLoginSuccess(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	vs := url.Values{"name":{"steve"}, "pw":{"massey"}}
	res, err := client.PostForm("http://localhost/api/login", vs)
	if err != nil {
		t.Error(err)
		return
	}
	if res.StatusCode != 200 {
		t.Error(err)
	}
}

func TestLoginFail(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	vs := url.Values{"name":{"test"}, "pw":{"nothing"}}
	res, err := client.PostForm("http://localhost/api/login", vs)
	if err != nil {
		t.Error(err)
		return
	}
	if res.StatusCode != 401 {
		t.Error(err)
	}
}

func TestLoginSuper(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	vs := url.Values{"name":{"steve"}, "pw":{"massey"}}
	res, err := client.PostForm("http://localhost/api/login", vs)
	if err != nil {
		t.Error(err)
		return
	}
	for _, cookie := range res.Cookies(){
		if cookie.Name == "super" && cookie.Value == "true" {
			return
		}
	}
	t.Fatal()
}

var cookieNormal []*http.Cookie
func TestLoginNormal(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	vs := url.Values{"name":{"test"}, "pw":{"abc"}}
	res, err := client.PostForm("http://localhost/api/login", vs)
	if err != nil {
		t.Error(err)
		return
	}
	for _, cookie := range res.Cookies(){
		if cookie.Name == "super" && cookie.Value == "true" {
			t.Fatal()
		}
	}
	cookieNormal = client.Jar.Cookies(u)
}

func login() {
	if cookieNormal == nil {
		vs := url.Values{"name":{"test"}, "pw":{"abc"}}
		_, _ = client.PostForm("http://localhost/api/login", vs)
	} else {
		client.Jar.SetCookies(u, cookieNormal)
	}
}

func TestGetPins(t *testing.T) {
	res, err := client.Get("http://localhost/api/getPins?north=-40&south=-41&east=176&west=175&timeBegin=0&timeEnd=20000")
	if err != nil {
		t.Error(err)
		return
	}
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if len(all) <= 100 {
		t.Error("response is too short")
	}
}

func TestGetPin(t *testing.T) {
	res, err := client.Get("http://localhost/api/getPin?information=17")
	if err != nil {
		t.Error(err)
		return
	}
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if len(all) <= 10 {
		t.Error("the response is too short")
	}
}

func TestStatic(t *testing.T) {
	res, err := client.Get("http://localhost")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(res.StatusCode)
	if res.ContentLength <= 100 {
		t.Error(err)
	}
}
