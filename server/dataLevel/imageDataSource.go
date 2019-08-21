package dataLevel

import (
	"github.com/ChenXingyuChina/asynchronousIO"
	"fmt"
	"os"
)

type imageDataSource struct {
	root string
}

func (i *imageDataSource) Load(key asynchronousIO.Key) (asynchronousIO.Bean, error) {
	f, err := os.Open(fmt.Sprintf(i.root, int64(key.(ImageKey))))
	if err != nil {
		return nil, err
	}
	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	length := uint64(s.Size())
	goal := GenImage(int64(key.(ImageKey)), length, false)
	_, err = f.Read(goal.data)
	if err != nil {
		return nil, err
	}
	err = f.Close()
	return goal, err
}

func (i *imageDataSource) Save(bean asynchronousIO.Bean) error {
	f, err := os.Create(fmt.Sprintf(i.root, bean.(*ImageResource).Id))
	if err != nil {
		return err
	}
	_, err = f.Write(bean.(*ImageResource).data)
	return err
}

func (i *imageDataSource) Delete(key asynchronousIO.Key) error {
	return os.Remove(fmt.Sprintf(i.root, int64(key.(ImageKey))))
}

