package actions

import (
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func GetMedia(w http.ResponseWriter, r *http.Request)  {
	vs := r.URL.Query()
	id, err := strconv.ParseInt(vs.Get("id"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	m, ok := buffer.MainCache.Load(base.MediaKey(id))
	if !ok {
		w.WriteHeader(404)
		return
	}
	media := m.(*base.Media)
	if media.Url == "file" {
		i, ok := buffer.MainCache.Load(dataLevel.ImageKey(id))
		if !ok {
			w.WriteHeader(404)
			return
		}
		image := i.(*dataLevel.ImageResource)
		http.SetCookie(w, &http.Cookie{Name:"type", Value:"f"})
		_, err = w.Write(image.Data)
	} else {
		http.SetCookie(w, &http.Cookie{Name:"type", Value:"u"})
		_, err = w.Write([]byte(media.Url))
	}
}
