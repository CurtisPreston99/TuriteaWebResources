package dataLevel

import (
	"TuriteaWebResources/asynchronousIO"
	"strconv"
)

type ImageResource struct {
	// todo create a pool for this type
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
