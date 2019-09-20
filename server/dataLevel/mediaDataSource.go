package dataLevel

import (
	"os"
	"strconv"
	"strings"

	"github.com/ChenXingyuChina/asynchronousIO"

	"TuriteaWebResources/server/base"
)

type mediaDataSource struct{}

func (mediaDataSource) Load(key asynchronousIO.Key) (asynchronousIO.Bean, error) {
	b := SQLWorker.GetMedia(int64(key.(base.MediaKey)))
	if b == nil {
		return nil, asynchronousIO.AsynchronousIOError{E: &os.PathError{Path: "media:" + strconv.FormatInt(int64(key.(base.MediaKey)), 16)}}
	}
	if strings.Compare(b.Url, "file") != 0 {
		OnLoadMedia(ImageKey(b.Uid))
	}
	return b, nil
}

func (mediaDataSource) Save(bean asynchronousIO.Bean) error {
	SQLWorker.ChangeMedia(bean.(*base.Media))
	return nil
}

func (mediaDataSource) Delete(key asynchronousIO.Key) error {
	return SQLWorker.DeleteMedia(int64(key.(base.MediaKey)))
}
