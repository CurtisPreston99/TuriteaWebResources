package actions

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func Update(w http.ResponseWriter, r *http.Request) {
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
		cid, err := strconv.ParseInt(r.Form.Get("id"), 16, 64)
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
		buffer.MainCache.Update(dataLevel.GenContentWithData(cid, resource, content))
	case dataLevel.ImagesResources:
		f, head, err := r.FormFile("data")
		err = r.ParseForm()
		if err != nil {
			return
		}
		title := r.Form.Get("title")
		if err != nil {
			w.WriteHeader(400)
			return
		}
		if err != nil {
			w.WriteHeader(400)
			return
		}
		b := make([]byte, head.Size)
		_, err = f.Read(b)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		id := base.GenMediaId()
		if !buffer.MainCache.CreateMedia(base.GenMedia(id, imageLocal, title, "file")) {
			w.WriteHeader(500)
			return
		}
		buffer.MainCache.CreateImage(b, id)
		_ = f.Close()
		err = r.MultipartForm.RemoveAll()
		if err != nil {
			fmt.Println(time.Now().Format(time.Stamp), err)
		}
		_, _ = w.Write([]byte("1"))
	case dataLevel.Pin:
		num, err := strconv.ParseInt(r.Form.Get("num"), 16, 64)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		data := r.Form.Get("data")
		jins, err := base.JsonToPins(strings.NewReader(data), uint16(num))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		for _, v := range jins {
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
}
