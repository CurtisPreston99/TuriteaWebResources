package base

import (
	"sync"
)

type Article struct {
	Id int64
	WriteBy int64
	Summary string
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
	articlePool.Put(articlePool)
	if delete {
		articleIdRecycle <- article.Id
	}
}

