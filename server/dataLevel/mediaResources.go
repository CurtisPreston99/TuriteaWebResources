package dataLevel

import (
	"TuriteaWebResources/asynchronousIO"
	"strconv"
	"sync"
)

type ImageResource struct {
	Id int64
	data []byte
}

func (i *ImageResource) GetKey() asynchronousIO.Key {
	return ImageKey(i.Id)
}

type ImageKey int64

func (i ImageKey) TypeId() int64 {
	return ImagesResources
}

func (i ImageKey) UniqueId() (int64, bool) {
	return int64(i), true
}

func (i ImageKey) ToString() (string, bool) {
	return strconv.FormatInt(int64(i), 16), true
}

var mediaResourcePool = new(sync.Pool)

var mediaIdChan = make(chan int64, 100)
var mediaIdRecycle = make(chan int64, 100)

func mediaIdProvider() {
	var id int64
	for {
		select {
		case mediaIdChan <- id:
			id++
		case i := <-mediaIdRecycle:
			mediaIdChan <- i
		}
	}
}

func init() {
	mediaResourcePool.New = func() interface{} {
		return &ImageResource{}
	}
	go mediaIdProvider()
}

func GenImage(id int64, dataLength uint64, newOne bool) *ImageResource {
	if newOne {
		id = <-mediaIdChan
	}
	goal := mediaResourcePool.Get().(*ImageResource)
	goal.Id = id
	goal.data = make([]byte, dataLength)
	return goal
}

func CreateImageByData(data []byte) *ImageResource {
	goal := mediaResourcePool.Get().(*ImageResource)
	goal.Id = <- mediaIdChan
	goal.data = data
	return goal
}

func RecycleImage(i *ImageResource, delete bool) {
	if delete {
		mediaIdRecycle <- i.Id
	}
	mediaResourcePool.Put(i)
}
