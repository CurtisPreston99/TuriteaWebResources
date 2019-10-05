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

func addArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("call add article")
	p, id := se.checkPermission(r)
	switch p {
	case super:
		fallthrough
	case normal:
		num, err := strconv.ParseInt(r.URL.Query().Get("num"), 16, 64)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = r.ParseForm()
		if err != nil {
			w.WriteHeader(400)
			return
		}
		data := r.Form.Get("data")
		articles, err := base.JsonToArticles(strings.NewReader(data), uint16(num))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		for _, v := range articles {
			v.Id = base.GenArticleId()
			v.WriteBy = id
			v.HomeContent = buffer.MainCache.CreateArticleContent([]dataLevel.Resource{}, "")
			if !buffer.MainCache.CreateArticle(v) {
				_, _ = fmt.Fprint(w, "f", " ")
				buffer.MainCache.Delete(dataLevel.ArticleContentKey(v.HomeContent))
				base.RecycleArticle(v, true)
			} else {
				_, _ = fmt.Fprint(w, strconv.FormatInt(v.Id, 16), strconv.FormatInt(v.HomeContent ,16), " ")
			}
		}

		makeCookie(w, id)
		se.renew(id)
	case public:
		w.WriteHeader(401)
	}
}


func addArticleWithImage(w http.ResponseWriter, r *http.Request) {
	log.Println("call add article with images")
	p, id := se.checkPermission(r)
	switch p {
	case super:
		fallthrough
	case normal:
		// pin[] images[] article[] fragment
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
		if len(images)|len(imageNum)|len(articles)|len(content) == 0 {
			w.WriteHeader(400)
			return
		}

		// parse data
		n, err := strconv.ParseInt(imageNum, 16, 64)
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

		// link new pin and new article
		v := allArticles[0]
		v.Id = base.GenArticleId()
		_, err = fmt.Fprint(w, v.Id)
		if err != nil {
			base.RecycleArticle(v, true)
			w.WriteHeader(400)
			return
		}
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
		v.WriteBy = id
		res := make([]dataLevel.Resource, n)
		for i := range res {
			res[i] = dataLevel.Resource{dataLevel.Image, imageIds[i].(int64)}
		}
		v.HomeContent = buffer.MainCache.CreateArticleContent(res, content)
		if !buffer.MainCache.CreateArticle(v) {
			buffer.MainCache.Delete(dataLevel.ArticleContentKey(v.HomeContent))
			base.RecycleArticle(v, true)
			w.WriteHeader(500)
			for _, id := range imageIds {
				buffer.MainCache.Delete(base.MediaKey(id.(int64)))
				buffer.MainCache.Delete(dataLevel.ImageKey(id.(int64)))
			}
			return
		}

	case public:
		w.WriteHeader(401)
	}

}
