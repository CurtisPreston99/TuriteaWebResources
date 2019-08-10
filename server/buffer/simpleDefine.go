package buffer

import (
	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/dataLevel"
)

//todo finish the buffer like this

func GetArticle(articleId int64, sqlLinker *dataLevel.SqlLinker) *base.Article {
	return sqlLinker.LoadArticle(articleId)
}

func GetArticleContent() {

}
