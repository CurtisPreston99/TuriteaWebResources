package actions

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func updateArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("call update article")
	p, id := se.checkPermission(r)
	switch p {
	case super:
		fallthrough
	case normal:
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(400)
			return
		}

		// get data
		images := r.Form.Get("images")
		imageNum := r.Form.Get("imageNum")
		articles := r.Form.Get("articles")
		content := r.Form.Get("content")
		aid := r.Form.Get("aid");
		if len(aid)|len(images)|len(imageNum)|len(articles)|len(content) == 0 {
			w.WriteHeader(400)
			return
		}

		// parse data
		articleId, err := strconv.ParseInt(aid, 16, 64)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		n, err := strconv.ParseInt(imageNum, 16, 64)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		allImage := make([]imageHelp, 0, n)
		err = json.Unmarshal([]byte(images), &allImage)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		allArticles, err := base.JsonToArticles(strings.NewReader(articles), 1)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		article := allArticles[0]
		article.Id = articleId


		imageIds := make([]interface{}, 0, n)
		for _, v := range allImage {
			b, err := base64.StdEncoding.DecodeString(v.Image)
			if err != nil {
				w.WriteHeader(400)
				for _, id := range imageIds {
					buffer.MainCache.Delete(base.MediaKey(id.(int64)))
					buffer.MainCache.Delete(dataLevel.ImageKey(id.(int64)))
				}
				return
			}
			id := base.GenMediaId()
			if !buffer.MainCache.CreateMedia(base.GenMedia(id, imageLocal, v.Title, "file")) {
				w.WriteHeader(500)
				for _, id := range imageIds {
					buffer.MainCache.Delete(base.MediaKey(id.(int64)))
					buffer.MainCache.Delete(dataLevel.ImageKey(id.(int64)))
				}
				return
			}
			buffer.MainCache.CreateImage(b, id)
			imageIds = append(imageIds, id)
		}

		content = fmt.Sprintf(content, imageIds...)
		// create article fragment and article
		article.WriteBy = id
		res := make([]dataLevel.Resource, n)
		for i := range res {
			res[i] = dataLevel.Resource{dataLevel.Image, imageIds[i].(int64)}
		}
		article.HomeContent = buffer.MainCache.CreateArticleContent(res, content)
		buffer.MainCache.Update(article)
	case public:
		w.WriteHeader(401)
	}
}

