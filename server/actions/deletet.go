package actions

import (
	"log"
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func deleteData(w http.ResponseWriter, r *http.Request) {
	log.Println("call delete")
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
	//fmt.Println(t)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	var ResId int64
	ResId, err = strconv.ParseInt(vs.Get("id"), 16, 64)
	//fmt.Println(information)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	switch t {
	case dataLevel.ImagesResources:
		buffer.MainCache.Delete(dataLevel.ImageKey(ResId))
	case dataLevel.ArticleContentResources:
		buffer.MainCache.Delete(dataLevel.ArticleContentKey(ResId))
	case dataLevel.Media:
		if !buffer.MainCache.Delete(base.MediaKey(ResId)) {
			w.WriteHeader(500)
			return
		}
		buffer.MainCache.Delete(dataLevel.ImageKey(ResId))
	case dataLevel.Article:
		if !buffer.MainCache.Delete(base.ArticleKey(ResId)) {
			w.WriteHeader(500)
			return
		}
	case dataLevel.Pin:
		if !buffer.MainCache.Delete(base.PinKey(ResId)) {
			w.WriteHeader(500)
			return
		}
	default:
		w.WriteHeader(400)
		return
	}
	_, _ = w.Write([]byte("ok"))
}
