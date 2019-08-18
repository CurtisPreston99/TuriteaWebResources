package base

import (
	"TuriteaWebResources/asynchronousIO"
	"sync"
)

type Article struct {
	Id int64
	WriteBy int64
	Summary string
}

func (a *Article) GetKey() asynchronousIO.Key {
	return ArticleKey(a.Id)
}

var articlePool *sync.Pool
var articleIdChan = make(chan int64, 100)
var articleIdRecycle = make(chan int64, 100)
func init() {
	articlePool = new(sync.Pool)
	articlePool.New = func() interface{} {
		return &Article{}
	}
	go articleIdProvider()
}

func articleIdProvider() {
	var id int64
	for {
		select {
		case articleIdChan<-id:
			id ++
		case Id := <-articleIdRecycle:
			articleIdChan <- Id
		}
	}
}

func GenArticle(Id, writeBy int64, summary string, newOne bool) *Article {
	if newOne{
		Id = <-articleIdChan
	}
	a := articlePool.Get().(*Article)
	a.Id = Id
	a.Summary = summary
	a.WriteBy = writeBy
	return a
}

func RecycleArticle(article *Article, delete bool) {
	if delete {
		articleIdRecycle <- article.Id
	}
	articlePool.Put(articlePool)
}

type ArticleKey int64

func (a ArticleKey) UniqueId() (int64, bool) {
	return int64(a), true
}

func (ArticleKey) ToString() (string, bool) {
	panic("")
}

func (ArticleKey) TypeId() int64 {
	return 4
}


