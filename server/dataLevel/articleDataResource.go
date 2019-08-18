package dataLevel

import (
	"TuriteaWebResources/asynchronousIO"
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

type articleContentDataSource struct {
	root string
	onLoadId func([]Resource)
}

func (a *articleContentDataSource) Load(key asynchronousIO.Key) (asynchronousIO.Bean, error) {
	f, err := os.Open(fmt.Sprintf(a.root, int64(key.(ArticleContentKey))))
	if err != nil {
		return nil, err
	}

	var length uint64
	err = binary.Read(f, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	var length2 uint64
	err = binary.Read(f, binary.LittleEndian, &length2)
	if err != nil {
		return nil, err
	}

	goal := GenContent(int64(key.(ArticleContentKey)), length, length2, false)
	var head reflect.SliceHeader
	head = *((*reflect.SliceHeader)(unsafe.Pointer(&(goal.resourcesId))))
	head.Cap = int(16 * length)
	head.Len = int(16 * length)
	_, err = f.Read(*(*[]byte)(unsafe.Pointer(&(head))))
	if err != nil {
		RecycleContent(goal, false)
		return nil, err
	}
	go a.onLoadId(goal.resourcesId)
	_, err = f.Read(goal.content)
	err = binary.Read(f, binary.LittleEndian, goal.content)
	if err != nil {
		RecycleContent(goal, false)
		return nil, err
	}
	return goal, nil
}

func (a *articleContentDataSource) Save(bean asynchronousIO.Bean) error {
	b := bean.(*ArticleResource)
	key := bean.GetKey()
	f, err := os.Create(fmt.Sprintf(a.root, int64(key.(ArticleContentKey))))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.LittleEndian, uint64(len(b.resourcesId)))
	if err != nil {
		return err
	}

	err = binary.Write(f, binary.LittleEndian, uint64(len(b.content)))
	if err != nil {
		return err
	}

	var t reflect.SliceHeader
	t = *(*reflect.SliceHeader)(unsafe.Pointer(&(b.resourcesId)))
	t.Len *= 16
	t.Cap *= 16
	err = binary.Write(f, binary.LittleEndian, *(*[]byte)(unsafe.Pointer(&t)))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.LittleEndian, b.content)
	return err
}

func (a *articleContentDataSource) Delete(key asynchronousIO.Key) error {
	return os.Remove(fmt.Sprintf(a.root, int64(key.(ArticleContentKey))))
}
