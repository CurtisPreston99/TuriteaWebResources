package dataLevel

import (
	"TuriteaWebResources/asynchronousIO"
	"TuriteaWebResources/server/base"
	"os"
	"strconv"
)

type pinDataSource struct {}

func (pinDataSource) Load(key asynchronousIO.Key) (asynchronousIO.Bean, error) {
	bean, ok := SQLWorker.GetPinById(int64(key.(base.PinKey)))
	if ok {
		return bean, nil
	}
	return nil, asynchronousIO.AsynchronousIOError{E:&os.PathError{Path:string(key.(base.PinKey))}}
}

func (pinDataSource) Save(bean asynchronousIO.Bean) error {
	pin := bean.(*base.Pin)
	SQLWorker.UpdatePin(pin)
	return nil
}

func (pinDataSource) Delete(key asynchronousIO.Key) error {
	ok := SQLWorker.DeletePin(int64(key.(base.PinKey)))
	if ok {
		// todo change other path error to this format
		return &asynchronousIO.AsynchronousIOError{E:&os.PathError{Path:"pin:" + strconv.FormatInt(int64(key.(base.PinKey)), 16)}}
	}
	return nil
}

