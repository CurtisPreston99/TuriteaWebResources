package actions

import (
	"log"
	"net/http"
	"strconv"

	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func getImageLocal(w http.ResponseWriter, r *http.Request) {
	log.Println("call get image local")
	//fmt.Println(r.URL)
	q := r.URL.Query()
	idSting := q.Get("id")
	if len(idSting) == 0 {
		w.WriteHeader(400)
		return
	}
	id, err := strconv.ParseInt(idSting, 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	b, ok := buffer.MainCache.Load(dataLevel.ImageKey(id))
	if !ok {
		w.WriteHeader(500)
		return
	}
	image := b.(*dataLevel.ImageResource)
	_, err = w.Write(image.Data)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
