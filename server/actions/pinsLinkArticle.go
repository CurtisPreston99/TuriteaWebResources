package actions

import (
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func PinsByArticle(w http.ResponseWriter, r *http.Request) {
	vs := r.URL.Query()
	id, err := strconv.ParseInt(vs.Get("id"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	pinIds := dataLevel.SQLWorker.SearchPinsIdWithArticle(id)
	if pinIds == nil {
		w.WriteHeader(400)
		return
	}
	pins := make([]*base.Pin, len(pinIds))
	for i, v := range pinIds {
		b, ok := buffer.MainCache.Load(base.PinKey(v))
		if !ok {
			pins[i] = nil
		} else {
			pins[i] = b.(*base.Pin)
		}
	}
	err = base.PinsToJson(pins, w)
	if err != nil {
		w.WriteHeader(400)
		return
	}
}

func ArticlesByPin(w http.ResponseWriter, r *http.Request) {
	vs := r.URL.Query()
	id, err := strconv.ParseInt(vs.Get("id"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	pinIds := dataLevel.SQLWorker.SelectArticlesIdWithPin(id)
	if pinIds == nil {
		w.WriteHeader(400)
		return
	}
	pins := make([]*base.Article, len(pinIds))
	for i, v := range pinIds {
		b, ok := buffer.MainCache.Load(base.ArticleKey(v))
		if !ok {
			pins[i] = nil
		} else {
			pins[i] = b.(*base.Article)
		}
	}
	err = base.ArticlesToJson(pins, w)
	if err != nil {
		w.WriteHeader(400)
		return
	}
}

func LinkArticleAndPin(w http.ResponseWriter, r *http.Request) {
	vs := r.URL.Query()
	aid, err := strconv.ParseInt(vs.Get("aid"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	pid, err := strconv.ParseInt(vs.Get("pid"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if dataLevel.SQLWorker.LinkPinToArticle(pid, aid) {
		w.WriteHeader(400)
	}
}
