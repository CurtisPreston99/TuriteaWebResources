package actions

import (
	"log"
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
)

func getMedia(w http.ResponseWriter, r *http.Request)  {
	log.Println("call get media")
	vs := r.URL.Query()
	id, err := strconv.ParseInt(vs.Get("information"), 16, 64)
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
	err = base.MediaToJson(media, w)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
