package actions

import (
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	p, uid := se.checkPermission(r)
	if p == public {
		w.WriteHeader(401)
		return
	} else {
		makeCookie(w, uid)
		se.renew(uid)
	}
	vs := r.URL.Query()
	t, err := strconv.ParseInt(vs.Get("type"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	id, err := strconv.ParseInt(vs.Get("id"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	switch t {
	case dataLevel.ImagesResources:
		if !buffer.MainCache.Delete(dataLevel.ImageKey(id)) {
			w.WriteHeader(500)
		}
	case dataLevel.ArticleContentResources:
		if !buffer.MainCache.Delete(dataLevel.ArticleContentKey(id)) {
			w.WriteHeader(500)
		}
	case dataLevel.Media:
		if !buffer.MainCache.Delete(base.MediaKey(id)) {
			w.WriteHeader(500)
		}
	case dataLevel.Article:
		if !buffer.MainCache.Delete(base.ArticleKey(id)) {
			w.WriteHeader(500)
		}
	case dataLevel.Pin:
		if !buffer.MainCache.Delete(base.PinKey(id)) {
			w.WriteHeader(500)
		}
	default:
		w.WriteHeader(400)
		return
	}
	_, _ = w.Write([]byte("-1"))
}
