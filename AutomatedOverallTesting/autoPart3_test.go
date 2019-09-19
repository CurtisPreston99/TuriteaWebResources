package AutomatedOverallTesting

import (
	"fmt"
	"net/url"
	"testing"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func TestGetMedia(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	res, err := client.Get("http://localhost/api/media?id=2")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
		return
	}
	if res.ContentLength < 20 {
		body := make([]byte, res.ContentLength)
		_, _ = res.Body.Read(body)
		t.Error(body)
	}
}

func TestGetFragment(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	buffer.MainCache.Load(base.MediaKey(2))
	buffer.MainCache.CreateArticleContent([]dataLevel.Resource{
		{Type:dataLevel.Image, Id:2},
		{Type:dataLevel.Image, Id:0},
		{Type:dataLevel.ArticleContent, Id:0},
	}, "abc")
	res, err := client.Get("http://localhost/api/fragment?id=2")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
		return
	}
	if res.ContentLength < 20 {
		body := make([]byte, res.ContentLength)
		_, _ = res.Body.Read(body)
		fmt.Println(string(body))
		t.Error(body)
	}
}

func TestGetImage(t *testing.T) {
	buffer.MainCache.CreateImage([]byte{123:1}, 0)
	client.Jar.SetCookies(u, cookieInit)
	res, err := client.Get("http://localhost/api/getImage?id=0")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
		return
	}
	if res.ContentLength < 20 {
		body := make([]byte, res.ContentLength)
		_, _ = res.Body.Read(body)
		fmt.Println(body)
		t.Error(body)
	}
}

func TestPinsByArticle(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	res, err := client.Get("http://localhost/api/pinsByArticle?id=2")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
		return
	}
	if res.ContentLength < 20 {
		body := make([]byte, res.ContentLength)
		_, _ = res.Body.Read(body)
		fmt.Println(body)
		t.Error(body)
	}
}

func TestArticlesByPin(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	res, err := client.Get("http://localhost/api/articlesByPin?id=1")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
		return
	}
	if res.ContentLength < 20 {
		body := make([]byte, res.ContentLength)
		_, _ = res.Body.Read(body)
		fmt.Println(body)
		t.Error(body)
	}
}

func TestLastArticle(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	res, err := client.Get("http://localhost/api/lastArticle")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
		return
	}
	if res.ContentLength < 20 {
		body := make([]byte, res.ContentLength)
		_, _ = res.Body.Read(body)
		fmt.Println(body)
		t.Error(body)
	}
}

func TestLastArticleWithArgument(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	res, err := client.Get("http://localhost/api/lastArticle?begin=1&num=2")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
		return
	}
	if res.ContentLength < 20 {
		body := make([]byte, res.ContentLength)
		_, _ = res.Body.Read(body)
		fmt.Println(body)
		t.Error(body)
	}
}

func TestSendFeedback(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	vs := url.Values{"email":{"123@123.com"}, "name":{"test"}, "feedback":{"abc"}}
	res, err := client.PostForm("http://localhost/api/sendfeedback", vs)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}
func TestGetArticlePage(t *testing.T) {
	client.Jar.SetCookies(u, cookieInit)
	res, err := client.Get("http://localhost/api/sendfeedback")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}