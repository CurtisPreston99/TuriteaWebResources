package dataLevel

import (
	"github.com/ChenXingyuChina/asynchronousIO"
	"github.com/ChenXingyuChina/asynchronousIO/M2N"
	"TuriteaWebResources/server/base"
	"runtime"
)

var OnLoadResourceId func([]Resource)
var OnLoadMedia func(key ImageKey)
var fileIO asynchronousIO.AsynchronousIOMachine
func Init() {
	dataSource := []asynchronousIO.DataSource{
		//&imageDataSource{"../../resources/temPictures/%x.img"},
		//&articleContentDataSource{root: "../../articles/%x.art", onLoadId: OnLoadResourceId,},
		&imageDataSource{"../../resources/temPictures/%x.img"},
		&articleContentDataSource{root: "../../articles/%x.art", onLoadId: OnLoadResourceId,},
		articleDataSource{},
		mediaDataSource{},
		pinDataSource{},
	}
	fileIO = M2N.NewM2NMachine(dataSource, 5, 32, false)
}

func LoadArticleContent(key ArticleContentKey) func() (asynchronousIO.Bean, error) {
	return fileIO.Load(key, 1)
}
func SaveArticleContent(resource *ArticleResource) {
	fileIO.Save(resource, 1)
}
func SaveArticleContentAndNotify(resource *ArticleResource) func() error {
	return fileIO.SaveAndCallBackWhenFinish(resource, 1)
}
func DeleteArticleContent(key ArticleContentKey) func() error {
	return fileIO.Delete(key, 1)
}

func LoadImage(key ImageKey) func() (asynchronousIO.Bean, error) {
	return fileIO.Load(key, 0)
}
func SaveImage(resource *ImageResource) {
	fileIO.Save(resource, 0)
}
func SaveImageAndNotify(resource *ImageResource) func() error {
	return fileIO.SaveAndCallBackWhenFinish(resource, 0)
}
func DeleteImage(key ImageKey) func() error {
	return fileIO.Delete(key, 0)
}

func LoadArticle(key base.ArticleKey) func() (asynchronousIO.Bean, error)  {
	return fileIO.Load(key, 2)
}
func SaveArticle(resource *base.Article) {
	fileIO.Save(resource, 2)
}
func SaveArticleAndNotify(resource *base.Article) func() error {
	return fileIO.SaveAndCallBackWhenFinish(resource, 2)
}
func DeleteArticle(key base.ArticleKey) func() error {
	return fileIO.Delete(key, 2)
}

func LoadMedia(key base.MediaKey) func() (asynchronousIO.Bean, error)  {
	return fileIO.Load(key, 3)
}
func SaveMedia(resource *base.Media) {
	fileIO.Save(resource, 3)
}
func SaveMediaAndNotify(resource *base.Media) func() error {
	return fileIO.SaveAndCallBackWhenFinish(resource, 3)
}
func DeleteMedia(key base.MediaKey) func() error {
	return fileIO.Delete(key, 3)
}

func LoadPin(key base.PinKey) func() (asynchronousIO.Bean, error)  {
	return fileIO.Load(key, 4)
}
func SavePin(resource *base.Pin) {
	fileIO.Save(resource, 4)
}
func SavePinAndNotify(resource *base.Pin) func() error {
	return fileIO.SaveAndCallBackWhenFinish(resource, 4)
}
func DeletePin(key base.PinKey) func() error {
	return fileIO.Delete(key, 4)
}

func Load(key asynchronousIO.Key) func() (asynchronousIO.Bean, error) {
	switch k := key.(type) {
	case base.MediaKey:
		return LoadMedia(k)
	case base.ArticleKey:
		return LoadArticle(k)
	case base.PinKey:
		return LoadPin(k)
	case ArticleContentKey:
		return LoadArticleContent(k)
	case ImageKey:
		return LoadImage(k)
	}
	panic("error type")
}

func SaveAndNotify(bean asynchronousIO.Bean) func() error {
	switch b := bean.(type) {
	case *base.Pin:
		return SavePinAndNotify(b)
	case *base.Article:
		return SaveArticleAndNotify(b)
	case *base.Media:
		return SaveMediaAndNotify(b)
	case *ArticleResource:
		return SaveArticleContentAndNotify(b)
	case *ImageResource:
		return SaveImageAndNotify(b)
	}
	panic("error type")
}

func Delete(key asynchronousIO.Key) func() error {
	switch k := key.(type) {
	case base.MediaKey:
		return DeleteMedia(k)
	case base.ArticleKey:
		return DeleteArticle(k)
	case base.PinKey:
		return DeletePin(k)
	case ArticleContentKey:
		return DeleteArticleContent(k)
	case ImageKey:
		return DeleteImage(k)
	}
	panic("error type")
}

func recycleData(bean asynchronousIO.Bean) {
	switch b := bean.(type) {
	case *base.Pin:
		base.RecyclePin(b, false)
	case *base.Article:
		base.RecycleArticle(b, false)
	case *base.Media:
		base.RecycleMedia(b, false)
	case *ArticleResource:
		RecycleContent(b, false)
	case *ImageResource:
		RecycleImage(b)
	}
}

func RecycleData(bean asynchronousIO.Bean) {
	runtime.SetFinalizer(bean, recycleData)
}