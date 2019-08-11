package dataLevel

import (
	"TuriteaWebResources/asynchronousIO"
	"strconv"
	"sync"
)

type ArticleKey int64

func (k ArticleKey) TypeId() int64 {
	return ArticleResources
}

func (k ArticleKey) UniqueId() (int64, bool) {
	return int64(k), true
}

func (k ArticleKey) ToString() (string, bool) {
	return strconv.FormatInt(int64(k), 16), true
}

type ArticleResource struct {
	// todo create a pool for this type
	Id int64
	content []byte
	resourcesId []Resource
}

func (a *ArticleResource) GetKey() asynchronousIO.Key {
	return ArticleKey(a.Id)
}

var contentPool = new(sync.Pool)

func init() {
	contentPool.New = func() interface{} {
		return &ArticleResource{}
	}
}
