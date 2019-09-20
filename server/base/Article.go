package base

import (
	"encoding/json"
	"github.com/ChenXingyuChina/asynchronousIO"
	"io"
	"sync"
)

type Article struct {
	Id          int64  `json:"id"`
	WriteBy     int64  `json:"wby"`
	Summary     string `json:"sum"`
	HomeContent int64  `json:"home"`
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
	var id int64 = 2
	for {
		select {
		case articleIdChan <- id:
			id++
		case Id := <-articleIdRecycle:
			articleIdChan <- Id
		}
	}
}

func GenArticle(Id, writeBy, homeContent int64, summary string) *Article {
	a := articlePool.Get().(*Article)
	a.Id = Id
	a.Summary = summary
	a.WriteBy = writeBy
	a.HomeContent = homeContent
	return a
}

func GenArticleId() int64 {
	return <-articleIdChan
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
	panic("implement me")
}

func (ArticleKey) TypeId() int64 {
	return 4
}

func ArticlesToJson(articles []*Article, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(articles)
}

func JsonToArticles(r io.Reader, num uint16) ([]*Article, error) {
	d := json.NewDecoder(r)
	d.UseNumber()
	goal := make([]*Article, num)
	for i, _ := range goal {
		goal[i] = articlePool.Get().(*Article)
	}
	err := d.Decode(&goal)
	if err != nil {
		return nil, err
	}
	return goal, nil
}

func JsonToArticle(r io.Reader) (*Article, error) {
	d := json.NewDecoder(r)
	d.UseNumber()
	goal := articlePool.Get().(*Article)
	err := d.Decode(&goal)
	if err != nil {
		return nil, err
	}
	return goal, nil
}
