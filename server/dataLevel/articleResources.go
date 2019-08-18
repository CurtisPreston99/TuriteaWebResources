package dataLevel

import (
	"TuriteaWebResources/asynchronousIO"
	"strconv"
	"sync"
)

type ArticleContentKey int64

func (k ArticleContentKey) TypeId() int64 {
	return ArticleContentResources
}

func (k ArticleContentKey) UniqueId() (int64, bool) {
	return int64(k), true
}

func (k ArticleContentKey) ToString() (string, bool) {
	return strconv.FormatInt(int64(k), 16), true
}

type ArticleResource struct {
	Id int64
	content []byte
	resourcesId []Resource
}

func (a *ArticleResource) GetKey() asynchronousIO.Key {
	return ArticleContentKey(a.Id)
}

var contentPool = new(sync.Pool)

func init() {
	contentPool.New = func() interface{} {
		return &ArticleResource{}
	}
	go contentIdProvider()
}

var contentIdChan = make(chan int64, 100)
var contentIdRecycle = make(chan int64, 100)

func contentIdProvider() {
	var id int64
	for {
		select {
		case contentIdChan <- id:
			id++
		case i := <- contentIdRecycle:
			contentIdChan <- i
		}
	}
}

func GenContent(id int64, resourcesLength uint64, contentLength uint64, newOne bool) *ArticleResource {
	if newOne {
		id = <-contentIdChan
	}
	goal := contentPool.Get().(*ArticleResource)
	goal.Id = id
	// todo maybe these slices can be recycle in some way
	goal.resourcesId = make([]Resource, resourcesLength)
	goal.content = make([]byte, contentLength)
	return goal
}

func CreateContentByData(resources []Resource, contentLength string) *ArticleResource {
	goal := contentPool.Get().(*ArticleResource)
	goal.Id = <-contentIdChan
	// todo maybe these slices can be recycle in some way
	goal.resourcesId = resources
	goal.content = []byte(contentLength)
	return goal
}

func RecycleContent(a *ArticleResource, delete bool) {
	if delete {
		contentIdRecycle <- a.Id
	}
	contentPool.Put(a)
}
