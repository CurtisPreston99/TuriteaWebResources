package dataLevel

import (
	"TuriteaWebResources/server/base"
	"github.com/ChenXingyuChina/asynchronousIO"
	"os"
	"strconv"
)

type articleDataSource struct{}

func (articleDataSource) Load(key asynchronousIO.Key) (asynchronousIO.Bean, error) {
	b := SQLWorker.LoadArticle(int64(key.(base.ArticleKey)))
	if b != nil {
		return b, nil
	}
	return nil, asynchronousIO.AsynchronousIOError{E: &os.PathError{Path: "article:" + strconv.FormatInt(int64(key.(base.ArticleKey)), 16)}}
}

func (articleDataSource) Save(bean asynchronousIO.Bean) error {
	SQLWorker.ChangeArticle(bean.(*base.Article))
	return nil
}

func (articleDataSource) Delete(key asynchronousIO.Key) error {
	return SQLWorker.DeleteArticle(int64(key.(base.ArticleKey)))
}
