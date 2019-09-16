package AutomatedOverallTesting

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)
var newUser = make(map[string]interface{})
func TestAddUser(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	vs := url.Values{"name":{"steve"}, "pw":{"massey"}}
	_, err := client.PostForm("http://localhost/api/login", vs)
	if err != nil {
		t.Fatal(err)
	}
	vs = url.Values{"name":{"bob"}, "role":{"1"}}
	res, err := client.PostForm("http://localhost/api/addUser", vs)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
		return
	}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&newUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChangePassword(t *testing.T) {
	if len(newUser) == 0 {
		TestAddUser(t)
	}
	client.Jar.SetCookies(u, cookieInit)
	pw := fmt.Sprintf("%x", md5.New().Sum([]byte(newUser["password"].(string))))
	vs := url.Values{"name":{newUser["name"].(string)}, "pw":{pw}}
	res, err := client.PostForm("http://localhost/api/login", vs)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
	vs = url.Values{"new":{"abc"}, "old":{pw}}
	res, err = client.PostForm("http://localhost/api/changePassword", vs)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	vs := url.Values{"name":{"steve"}, "pw":{"massey"}}
	res, err := client.PostForm("http://localhost/api/login", vs)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatal("login super user fail: ", res.StatusCode)
	}
	vs = url.Values{"name":{"bob"}}
	res, err = client.PostForm("http://localhost/api/deleteUser", vs)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error("delete fail: ", res.StatusCode)
	}
}

func TestUnlink(t *testing.T) {
	login()
	res, err := client.Get("http://localhost/api/unlinkArticleAndPin?aid=2&pid=1")
	if err != nil {
		t.Fatal(err)
		return
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestDeleteMedia(t *testing.T) {
	login()
	res, err := client.Get("http://localhost/api/delete?type=3&id=2")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestDeleteArticle(t *testing.T) {
	login()
	res, err := client.Get("http://localhost/api/delete?type=4&id=2")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestDeletePin(t *testing.T) {
	login()
	res, err := client.Get("http://localhost/api/delete?type=2&id=1b")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestDeleteFragment(t *testing.T) {
	login()
	res, err := client.Get("http://localhost/api/delete?type=0&id=1")
	if err != nil {
		t.Fatal(err)
		return
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}
