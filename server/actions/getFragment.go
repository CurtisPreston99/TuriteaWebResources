package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

type fragmentHelper struct {
	Position int64 `json:"pos"`
	Type string `json:"type"`
	Id int64 `json:"id"`
	Media *base.Media `json:"m"`
}

type helpFragment struct {
	Content string `json:"content"`
	Res fragmentHelper `json:"res"`
}

func getFragment(w http.ResponseWriter, r *http.Request) {
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
		return
	}
	ac := b.(*dataLevel.ArticleResource)
	_, err = w.Write([]byte("{\"content\":"))
	if err != nil {
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(string(ac.Content))
	if err != nil {
		return
	}
	_, err = w.Write([]byte(", \"res\": ["))
	var start = true
	var helper = &fragmentHelper{}
	for i, v := range ac.ResourcesId {
		if !start{
			_, err = w.Write([]byte(","))
		}
		if v.Type & (dataLevel.Image | dataLevel.Video) != 0 {
			b, exist := buffer.MainCache.LoadIfExist(base.MediaKey(v.Id))
			if exist == buffer.Exist {
				media := b.(*base.Media)
				helper.Position = int64(i)
				helper.Type = "m"
				helper.Media = media
				helper.Id = 0
				err = encoder.Encode(helper)
				start = false
				if err != nil {
					w.WriteHeader(500)
					return
				}
			} else if exist == buffer.NotInBuffer {
				helper.Media = nil
				helper.Id = v.Id
				helper.Type = "mId"
				helper.Position = int64(i)
				start = false
				err = encoder.Encode(helper)
				if err != nil {
					w.WriteHeader(500)
					return
				}
			}
		} else {
			helper.Media = nil
			helper.Id = v.Id
			helper.Type = "f"
			helper.Position = int64(i)
			start = false
			err := encoder.Encode(helper)
			if err != nil {
				w.WriteHeader(500)
				return
			}
			start = false
			if err != nil {
				return
			}
		}
	}
	_, err = w.Write([]byte(fmt.Sprintf("], \"id\":%x}", ac.Id)))
}
