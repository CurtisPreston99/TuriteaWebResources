package dataLevel

import (
	"TuriteaWebResources/asynchronousIO"
	"encoding/binary"
	"fmt"
	"os"
)

type articleDataResource struct {
	root string
	onLoadId *func([]int64)
}

func (a *articleDataResource) Load(key asynchronousIO.Key) (asynchronousIO.Bean, error) {
	f, err := os.Open(fmt.Sprintf(a.root, int64(key.(ArticleKey))))
	if err != nil {
		return nil, err
	}
	var length uint64
	err = binary.Read(f, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	goal := &ArticleResource{Id:int64(key.(ArticleKey))}
	goal.resourcesId = make([]int64, length)
	err = binary.Read(f, binary.LittleEndian, goal.resourcesId)
	if err != nil {
		return nil, err
	}
	go (*a.onLoadId)(goal.resourcesId)
	err = binary.Read(f, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	goal.content = make([]byte, length)
	err = binary.Read(f, binary.LittleEndian, goal.content)
	return goal, err
}

func (a *articleDataResource) Save(bean asynchronousIO.Bean) error {
	b := bean.(*ArticleResource)
	key := bean.GetKey()
	f, err := os.Create(fmt.Sprintf(a.root, int64(key.(ArticleKey))))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.LittleEndian, uint64(len(b.resourcesId)))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.LittleEndian, b.resourcesId)
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.LittleEndian, uint64(len(b.resourcesId)))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.LittleEndian, b.content)
	return err
}

func (a *articleDataResource) Delete(key asynchronousIO.Key) error {
	return os.Remove(fmt.Sprintf(a.root, int64(key.(ArticleKey))))
}
