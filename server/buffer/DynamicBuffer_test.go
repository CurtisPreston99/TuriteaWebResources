package buffer

import (
	"fmt"
	"testing"
	"time"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/dataLevel"
)

func init() {
	dataLevel.OnLoadResourceId = func(resources []dataLevel.Resource) {

	}
	dataLevel.Init()
}

func TestCache_LoadNoExist(t *testing.T) {
	b, ok := MainCache.Load(dataLevel.ArticleContentKey(1))
	if !ok {
		t.Fatal(b)
	}
	//fmt.Println(b)
}

func TestCache_LoadExist(t *testing.T) {
	b, ok := MainCache.Load(dataLevel.ArticleContentKey(1))
	if !ok {
		t.Fatal(b)
	}
	b2, ok := MainCache.Load(dataLevel.ArticleContentKey(1))
	if !ok || b2 != b {
		t.Fatal(b, b2)
	}
}

func TestCache_UpdateNotExist(t *testing.T) {
	MainCache.Update(&base.Pin{Uid:1, Latitude:100, Longitude:-40, Description:"test3", TagType:"hospital", Time:18711, Name:"test", Owner:0})
	MainCache.flushBlock(1 + uint8(dataLevel.Pin))
	<-time.Tick(2*time.Second)
}

func TestCache_UpdateExist(t *testing.T) {
	MainCache.Update(&base.Pin{Uid:1, Latitude:100, Longitude:-40, Description:"test", TagType:"hospital", Time:18711, Name:"test", Owner:0})
	MainCache.Update(&base.Pin{Uid:1, Latitude:100, Longitude:-40, Description:"test2", TagType:"hospital", Time:18711, Name:"test", Owner:0})
	MainCache.flushBlock(1 + uint8(dataLevel.Pin))
	<-time.Tick(2*time.Second)
}

func TestCache_DeleteNotExist(t *testing.T) {
	if !MainCache.Delete(base.ArticleKey(2)) {
		t.Fatal()
	}
}

func TestCache_DeleteExist(t *testing.T) {
	article := &base.Article{base.GenArticleId(), 0, "233"}
	k := MainCache.CreateArticle(article)
	if !k {
		t.Fatal()
	}
	// uncomment it to check by hand
	//<- time.Tick(30*time.Second)
	if !MainCache.Delete(base.ArticleKey(2)) {
		t.Fatal()
	}
}

func TestCache_CreateArticle(t *testing.T) {
	article := &base.Article{base.GenArticleId(), 0, "233"}
	fmt.Println(article)
	if !MainCache.CreateArticle(article) {
		t.Fatal()
	}
}

func TestCache_CreateImage(t *testing.T) {
	k := base.GenMediaId()
	MainCache.CreateImage([]byte{123:1}, k)
	MainCache.flushBlock(uint8(k) + uint8(dataLevel.ImagesResources))
	<-time.Tick(2 * time.Second)
}

func TestCache_CreateArticleContent(t *testing.T) {
	k := MainCache.CreateArticleContent([]dataLevel.Resource{{1, 3}, {1, 5}}, "test")
	MainCache.flushBlock(uint8(k) + uint8(dataLevel.ArticleContentResources))
	<-time.Tick(2*time.Second)
}

func TestCache_CreateMedia(t *testing.T) {
	if !MainCache.CreateMedia(base.GenMedia(base.GenMediaId(), 1, "test", "htp")) {
		t.Fatal()
	}
}

func TestCache_CreatePin(t *testing.T) {
	if !MainCache.CreatePin(&base.Pin{base.GenPinId(), 0, 1, 1, 1, "jkss", "hospital", "testInBuffer", "#000011"}) {
		t.Fatal()
	}
}

func TestCache_LoadAsynchronousNotExist(t *testing.T) {
	MainCache.LoadAsynchronous(dataLevel.ArticleContentKey(1))
	<-time.Tick(1 * time.Second)
	b, ok := MainCache.Load(dataLevel.ArticleContentKey(1))
	if !ok {
		t.Fatal()
	}
	fmt.Println(b)
}

func TestCache_LoadAsynchronousExist(t *testing.T) {
	b, ok := MainCache.Load(dataLevel.ArticleContentKey(1))
	<-time.Tick(1 * time.Second)
	if !ok {
		t.Fatal()
	}
	MainCache.LoadAsynchronous(dataLevel.ArticleContentKey(1))
	<-time.Tick(1 * time.Second)
	b2, ok := MainCache.Load(dataLevel.ArticleContentKey(1))
	if !ok || b != b2{
		t.Fatal(b, b2)
	}
}
var lastKey int64
func TestLoadAfterCreate(t *testing.T) {
	media := base.GenMedia(base.GenMediaId(), 1, "test", "htp")
	k := MainCache.CreateMedia(media)
	lastKey = media.Uid
	if !k {
		t.Fatal()
	}
	b, ok := MainCache.Load(base.MediaKey(lastKey))
	if !ok {
		t.Fatal()
	}
	fmt.Println(b)
}

func TestLoadAfterDelete(t *testing.T) {
	if !MainCache.Delete(base.MediaKey(lastKey)) {
		t.Fatal()
	}

	_, ok := MainCache.Load(base.MediaKey(lastKey))
	if ok {
		t.Fatal()
	}
}

