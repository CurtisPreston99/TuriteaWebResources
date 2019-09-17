package buffer

import (
	"fmt"
	"testing"
	"time"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/dataLevel"
)

func TestMain(m *testing.M) {
	dataLevel.OnLoadResourceId = func(resources []dataLevel.Resource) {

	}
	dataLevel.OnLoadMedia = func(key dataLevel.ImageKey) {

	}
	dataLevel.Init()
	m.Run()
	MainCache.Delete(base.ArticleKey(100))
}

func TestCache_LoadNoExist(t *testing.T) {
	b, ok := MainCache.Load(base.PinKey(3))
	if !ok {
		t.Fatal(b)
	}
	//fmt.Println(b)
}

func TestCache_LoadExist(t *testing.T) {
	b, ok := MainCache.Load(base.PinKey(1))
	if !ok {
		t.Fatal(b)
	}
	b2, ok := MainCache.Load(base.PinKey(1))
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

func TestCache_CreateArticle(t *testing.T) {
	article := &base.Article{100, 0, "233", 1}
	fmt.Println(article)
	if !MainCache.CreateArticle(article) {
		t.Fatal()
	}
}

func TestCache_CreateImage(t *testing.T) {
	MainCache.CreateImage([]byte{123:1}, 100)
	MainCache.flushBlock(uint8(100) + uint8(dataLevel.ImagesResources))
	<-time.Tick(2 * time.Second)
}

var acId int64
func TestCache_CreateArticleContent(t *testing.T) {
	acId = MainCache.CreateArticleContent([]dataLevel.Resource{{1, 3}, {1, 5}}, "test")
	MainCache.flushBlock(uint8(acId) + uint8(dataLevel.ArticleContentResources))
	<-time.Tick(2*time.Second)
}

//func TestCache_CreateMedia(t *testing.T) {
//	if !MainCache.CreateMedia(base.GenMedia(100, 1, "test", "htp")) {
//		t.Fatal()
//	}
//}

func TestCache_CreatePin(t *testing.T) {
	if !MainCache.CreatePin(&base.Pin{100, 0, 1, 1, 1, "jkss", "hospital", "testInBuffer", "#000011"}) {
		t.Fatal()
	}
}

func TestCache_DeleteNotExist(t *testing.T) {
	if !MainCache.Delete(base.PinKey(100)) {
		t.Fatal()
	}
}

func TestCache_LoadAsynchronousNotExist(t *testing.T) {
	MainCache.LoadAsynchronous(base.PinKey(17))
	<-time.Tick(1 * time.Second)
	b, ok := MainCache.Load(base.PinKey(17))
	if !ok {
		t.Fatal()
	}
	fmt.Println(b)
}

func TestCache_LoadAsynchronousExist(t *testing.T) {
	b, ok := MainCache.Load(base.PinKey(18))
	<-time.Tick(1 * time.Second)
	if !ok {
		t.Fatal()
	}
	MainCache.LoadAsynchronous(base.PinKey(18))
	<-time.Tick(1 * time.Second)
	b2, ok := MainCache.Load(base.PinKey(18))
	if !ok || b != b2{
		t.Fatal(b, b2)
	}
}

func TestLoadAfterCreate(t *testing.T) {
	media := base.GenMedia(100, 1, "test", "htp")
	k := MainCache.CreateMedia(media)
	if !k {
		t.Fatal()
	}
	b, ok := MainCache.Load(base.MediaKey(100))
	if !ok {
		t.Fatal()
	}
	fmt.Println(b)
}

func TestCache_DeleteExist(t *testing.T) {
	if !MainCache.Delete(base.MediaKey(100)) {
		t.Fatal()
	}
}

func TestLoadAfterDelete(t *testing.T) {
	if !MainCache.Delete(base.MediaKey(100)) {
		t.Fatal()
	}

	_, ok := MainCache.Load(base.MediaKey(100))
	if ok {
		t.Fatal()
	}
}
