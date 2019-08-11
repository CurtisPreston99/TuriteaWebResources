package dataLevel

import (
	"TuriteaWebResources/asynchronousIO"
	"TuriteaWebResources/asynchronousIO/M2N"
)
// todo add more data sources and manage the data in database.
var onLoadResourceId func([]Resource) // todo init it first and dont change it when running
var fileIO asynchronousIO.AsynchronousIOMachine
func Init() {
	dataSource := []asynchronousIO.DataSource{&imageDataSource{"./resources/temPictures/%x.img"},
		&articleDataResource{root:"./articles/%x.art", onLoadId: onLoadResourceId}}
	fileIO = M2N.NewM2NMachine(dataSource, 2, 16, false)
}

func LoadArticleContent(key ArticleKey) func() (asynchronousIO.Bean, error) {
	return fileIO.Load(key, 1)
}

func SaveArticleContent(resource *ArticleResource) {
	fileIO.Save(resource, 1)
}

func SaveArticleContentAndNotify(resource *ArticleResource) func() error {
	return fileIO.SaveAndCallBackWhenFinish(resource, 1)
}

func DeleteArticleContent(key ArticleKey) func() error {
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
