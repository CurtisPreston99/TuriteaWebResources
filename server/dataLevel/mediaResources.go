package dataLevel

import (
	"github.com/ChenXingyuChina/asynchronousIO"
	"strconv"
	"sync"
)

type ImageResource struct {
	Id   int64
	Data []byte
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

func init() {
	mediaResourcePool.New = func() interface{} {
		return &ImageResource{}
	}
}

func GenImage(id int64, dataLength uint64) *ImageResource {
	goal := mediaResourcePool.Get().(*ImageResource)
	goal.Id = id
	goal.Data = make([]byte, dataLength)
	return goal
}

func CreateImageByData(data []byte) *ImageResource {
	goal := mediaResourcePool.Get().(*ImageResource)
	goal.Data = data
	return goal
}

func RecycleImage(i *ImageResource) {
	mediaResourcePool.Put(i)
}
