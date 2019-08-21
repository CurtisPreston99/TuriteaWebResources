package base

import (
	"github.com/ChenXingyuChina/asynchronousIO"
	"sync"
)

type Media struct {
	Type uint8
	Title string
	Url string
	Uid int64
}

func (m *Media) GetKey() asynchronousIO.Key {
	return MediaKey(m.Uid)
}

var mediaPool *sync.Pool
var mediaIdChan = make(chan int64, 100)
var mediaIdRecycle = make(chan int64, 100)
func init() {
	mediaPool = new(sync.Pool)
	mediaPool.New = func() interface{} {
		return &Media{}
	}
	go mediaIdProvider()
}

func mediaIdProvider() {
	var id int64 = 2
	for {
		select {
		case mediaIdChan <- id:
			id ++
		case Id := <-mediaIdRecycle:
			mediaIdChan <- Id
		}
	}
}

func GenMedia(uid int64, t uint8, title string, url string, NewOne bool) *Media {
	if NewOne {
		uid = <- mediaIdChan
	}
	media := mediaPool.Get().(*Media)
	media.Uid = uid
	media.Type = t
	media.Title = title
	media.Url = url
	return media
}

func RecycleMedia(media *Media, delete bool) {
	if delete {
		mediaIdRecycle <- media.Uid
	}
	mediaPool.Put(media)
}

type MediaKey int64

func (m MediaKey) UniqueId() (int64, bool) {
	return int64(m), true
}

func (MediaKey) ToString() (string, bool) {
	panic("implement me")
}

func (MediaKey) TypeId() int64 {
	return 3
}



