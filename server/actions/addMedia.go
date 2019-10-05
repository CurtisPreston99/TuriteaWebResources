package actions

import (
	"encoding/base64"
	"log"
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
)

const (
	imageLocal = 1 << iota
	video
	imageCloud
	articleContent
)

func addImage(w http.ResponseWriter, r *http.Request) {
	log.Println("call add image")
	p, uid := se.checkPermission(r)
	switch p {
	case normal:
		fallthrough
	case super:
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(400)
			return
		}
		data := r.Form.Get("data")
		title := r.Form.Get("title")
		if len(data) & len(title) == 0 {
			w.WriteHeader(400)
			return
		}
		b, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		id := base.GenMediaId()
		if !buffer.MainCache.CreateMedia(base.GenMedia(id, imageLocal, title, "file")) {
			w.WriteHeader(500)
			base.RecycleMedia(&base.Media{Uid:id}, true)
			return
		}
		buffer.MainCache.CreateImage(b, id)
		//_, _ = w.Write([]byte(`<html><body><p information="imageId">`))
		_, _ = w.Write([]byte(strconv.FormatInt(id, 16)))
		//_, _ = w.Write([]byte(`</p></body></html>`))
		makeCookie(w, uid)
		se.renew(uid)
	case public:
		w.WriteHeader(401)
	}
}
