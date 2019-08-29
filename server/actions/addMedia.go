package actions

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
)

const (
	imageLocal = 1 << iota
	video
	imageCloud
	articleContent
)

func AddImage(w http.ResponseWriter, r *http.Request) {
	log.Println("call add image")
	p, uid := se.checkPermission(r)
	switch p {
	case normal:
		fallthrough
	case super:
		f, head, err := r.FormFile("data")
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = r.ParseForm()
		if err != nil {
			return
		}
		title := r.Form.Get("title")
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
		makeCookie(w, uid)
		se.renew(uid)
	case public:
		w.WriteHeader(401)
	}
}
