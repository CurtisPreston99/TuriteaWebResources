package actions

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func pinsByArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("call pins by article")
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

func articlesByPin(w http.ResponseWriter, r *http.Request) {
	log .Println("call articles by pin")
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
		fmt.Println(v)
		b, ok := buffer.MainCache.Load(base.ArticleKey(v))
		//fmt.Println(b)
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

func linkArticleAndPin(w http.ResponseWriter, r *http.Request) {
	p, id := se.checkPermission(r)
	if p == public {
		w.WriteHeader(401)
	} else {
		se.renew(id)
		makeCookie(w, id)
	}
	log.Println("call link article and pin")
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
	//fmt.Println(pid, aid)
	if !dataLevel.SQLWorker.LinkPinToArticle(pid, aid) {
		w.WriteHeader(400)
	}
	_, _ = w.Write([]byte("ok"))
}
