package dataLevel

type Resource struct {
	Type uint8
	Id int64
}

const (
	ArticleContent = iota
	MediaResource
)
