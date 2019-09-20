package dataLevel

import (
	"TuriteaWebResources/server/base"
	"github.com/ChenXingyuChina/asynchronousIO"
	"os"
	"strconv"
)

type pinDataSource struct{}

func (pinDataSource) Load(key asynchronousIO.Key) (asynchronousIO.Bean, error) {
	bean, ok := SQLWorker.GetPinById(int64(key.(base.PinKey)))
	if ok {
		return bean, nil
	}
	return nil, asynchronousIO.AsynchronousIOError{E: &os.PathError{Path: string(key.(base.PinKey))}}
}

func (pinDataSource) Save(bean asynchronousIO.Bean) error {
	pin := bean.(*base.Pin)
	if !SQLWorker.UpdatePin(pin) {
		return asynchronousIO.AsynchronousIOError{E: &os.PathError{Path: "pin:" + strconv.FormatInt(int64(pin.Uid), 16)}}
	}
	return nil
}

func (pinDataSource) Delete(key asynchronousIO.Key) error {
	ok := SQLWorker.DeletePin(int64(key.(base.PinKey)))
	if !ok {
		return asynchronousIO.AsynchronousIOError{E: &os.PathError{Path: "pin:" + strconv.FormatInt(int64(key.(base.PinKey)), 16)}}
	}
	return nil
}
