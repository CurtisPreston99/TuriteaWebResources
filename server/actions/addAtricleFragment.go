package actions

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func addArticleFragment(w http.ResponseWriter, r *http.Request) {
	log.Println("call add article fragment")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	res := r.Form.Get("resIds")
	if len(res) == 0 {
		w.WriteHeader(400)
		return
	}
	content := r.Form.Get("content")
	if len(content) == 0 {
		w.WriteHeader(400)
		return
	}
	num, err := strconv.ParseInt(r.Form.Get("num"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	var bufferId []dataLevel.Resource
	if num != 0 {
		bufferId, err = dataLevel.JsonToResource(strings.NewReader(content), uint16(num))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		buffer.MainCache.CreateArticleContent(bufferId, content)
	} else {
		buffer.MainCache.CreateArticleContent([]dataLevel.Resource{}, content)
	}
	_, _ = w.Write([]byte("ok"))
}
