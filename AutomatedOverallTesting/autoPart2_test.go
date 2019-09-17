package AutomatedOverallTesting

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"testing"
	"time"
)

func TestAddPins(t *testing.T) {
	login()
	vs := url.Values{"data":{`[
{
	"owner": 0,
    "lat": 10.1,
    "lon": 10.2,
    "time": 55,
    "description": "abc",
    "tag_type": "airport",
    "name": "test1",
    "color": "#00ff00"
}]
`}}
	res, err := client.PostForm("http://localhost/api/addPins?num=1", vs)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestAddArticle(t *testing.T) {
	login()
	vs := url.Values{"data":{`[{
    "sum":"abc",
    "home":0
  }]`}}
	res, err := client.PostForm("http://localhost/api/addArticle?num=1", vs)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestAddArticleFragment(t *testing.T) {
	login()
	vs := url.Values{"resIds":{`[
  {"t":0, "id":1},
  {"t":1, "id":2}
]`}, "num":{"2"}, "content":{"abc"}}
	res, err := client.PostForm("http://localhost/api/addArticleFragment", vs)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestAddImage(t *testing.T) {
	login()
	buffer := &bytes.Buffer{}
	sendWrite := multipart.NewWriter(buffer)
	w, err := sendWrite.CreateFormFile("data", "a")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(w, bytes.NewReader([]byte("123432")))
	if err != nil {
		t.Error(err)
		return
	}
	err = sendWrite.WriteField("title", "ab")
	if err != nil {
		t.Error(err)
		return
	}
	err = sendWrite.Close()
	if err != nil {
		t.Error(err)
		return
	}
	res, err := client.Post("http://localhost/api/addImage", sendWrite.FormDataContentType(), buffer)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestUpdateFragment(t *testing.T) {
	login()
	vs := url.Values{"resIds":{`[{"t":0, "id":1},{"t":1, "id":2}]`},"num":{"2"},"content":{"absd"},"id":{"1"}}
	res, err := client.PostForm("http://localhost/api/update?type=0", vs)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestUpdatePin(t *testing.T) {
	login()
	vs := url.Values{"data":{`[
{
	"id":27,
	"owner": 0,
    "lat": 10.1,
    "lon": 10.2,
    "time": 55,
    "description": "abcd",
    "tag_type": "airport",
    "name": "test1",
    "color": "#00ff00"
}]
`}, "num":{"1"}}
	res, err := client.PostForm("http://localhost/api/update?type=2", vs)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}

func TestUpdateArticle(t *testing.T) {
	log.Println(time.Now())
	login()
	vs := url.Values{"data": {`{"id":2,
    "wby":1,
    "sum":"test",
    "home":3}`}}
	res, err := client.PostForm("http://localhost/api/update?type=4", vs)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
	log.Println(time.Now())
}

func TestLink(t *testing.T) {
	login()
	res, err := client.Get("http://localhost/api/linkPinToArticle?aid=2&pid=1")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}
