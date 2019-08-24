package actions

import (
	"log"
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func GetFragment (w http.ResponseWriter, r *http.Request) {
	log.Println("call get fragment")
	vs := r.URL.Query()
	id, err := strconv.ParseInt(vs.Get("id"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	b, ok := buffer.MainCache.Load(dataLevel.ArticleContentKey(id))
	if !ok {
		w.WriteHeader(404)
	}
	ac := b.(*dataLevel.ArticleResource)
	_, err = w.Write([]byte("{content:\""))
	if err != nil {
		return
	}
	_, err = w.Write(ac.Content)
	if err != nil {
		return
	}
	_, err = w.Write([]byte("\", resources: ["))
	var start bool
	for i, v := range ac.ResourcesId {
		if v.Type & (imageCloud | video) != 0 {
			if !start{
				_, err = w.Write([]byte(","))
			}
			b, exist := buffer.MainCache.LoadIfExist(base.MediaKey(v.Id))
			if exist == buffer.Exist {
				_, err = w.Write([]byte("{pos:"))
				if err != nil {
					return
				}
				_, err = w.Write([]byte(strconv.FormatInt(int64(i), 16)))
				_, err = w.Write([]byte(",type: \"u\", url:\""))
				if err != nil {
					return
				}
				_, err = w.Write([]byte(b.(*base.Media).Url))
				if err != nil {
					return
				}
				_, err = w.Write([]byte("\"}"))
				if err != nil {
					return
				}
			}
		}
	}
	_, err = w.Write([]byte("]}"))
}
