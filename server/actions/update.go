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

func update(w http.ResponseWriter, r *http.Request) {
	log.Println("call update")
	p, id := se.checkPermission(r)
	if p == public {
		w.WriteHeader(401)
		return
	} else {
		makeCookie(w, id)
		se.renew(id)
	}
	t, err := strconv.ParseInt(r.URL.Query().Get("type"), 16, 64)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
		return
	}
	switch t {
	case dataLevel.ArticleContentResources:
		num, err := strconv.ParseInt(r.Form.Get("num"), 16, 64)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		cid, err := strconv.ParseInt(r.Form.Get("information"), 16, 64)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		resource, err := dataLevel.JsonToResource(strings.NewReader(r.Form.Get("resIds")), uint16(num))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		content := r.Form.Get("content")
		if len(content) == 0 {
			w.WriteHeader(400)
			return
		}
		buffer.MainCache.Update(dataLevel.GenContentWithData(cid, resource, content))
	case dataLevel.Pin:
		num, err := strconv.ParseInt(r.Form.Get("num"), 16, 64)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		data := r.Form.Get("data")
		pins, err := base.JsonToPins(strings.NewReader(data), uint16(num))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		for _, v := range pins {
			buffer.MainCache.Update(v)
		}
	case dataLevel.Article:
		data := r.Form.Get("data")
		a, err := base.JsonToArticle(strings.NewReader(data))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		buffer.MainCache.Update(a)
	default:
		w.WriteHeader(400)
	}
	_, _ = w.Write([]byte("ok"))
}
