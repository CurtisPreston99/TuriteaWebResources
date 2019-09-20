package actions

import (
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
		state := true
		for i, v := range articles {
			v.Id = base.GenArticleId()
			v.WriteBy = id
			v.HomeContent = buffer.MainCache.CreateArticleContent([]dataLevel.Resource{}, "")
			if !buffer.MainCache.CreateArticle(v) {
				_, _ = fmt.Fprint(w, i, " ")
				buffer.MainCache.Delete(dataLevel.ArticleContentKey(v.HomeContent))
				base.RecycleArticle(v, true)
				state = false
			}
		}
		if !state {
			_, _ = w.Write([]byte("-1"))
		} else {
			_, _ = w.Write([]byte("ok"))
		}
		makeCookie(w, id)
		se.renew(id)
	case public:
		w.WriteHeader(401)
	}
}
