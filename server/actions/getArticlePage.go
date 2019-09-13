package actions

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func getArticlePage(w http.ResponseWriter, r *http.Request) {
	log.Println("get article")
	//todo send the template
	//hw := bufio.NewWriter(w)
	articleId := strings.SplitN(r.URL.Path, "/", 3)[3]
	if len(articleId) == 0 {
		w.WriteHeader(400)
		// fixme redirect to 404 page
		return
	}
	id, err := strconv.ParseInt(articleId, 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	b, ok := buffer.MainCache.Load(base.ArticleKey(id))
	if !ok {
		w.WriteHeader(404)
		return
	}
	home := b.(*base.Article).HomeContent
	http.SetCookie(w, &http.Cookie{Name:"home", Value:strconv.FormatInt(home, 16)})
	buffer.MainCache.LoadAsynchronous(dataLevel.ArticleContentKey(home))
}
