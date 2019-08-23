package actions

import (
	"net/http"
	"strconv"
	"strings"

	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func AddArticleFragment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
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
	bufferId, err := dataLevel.JsonToResource(strings.NewReader(content), uint16(num))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if -1 == buffer.MainCache.CreateArticleContent(bufferId, content){
		w.WriteHeader(500)
	}
}
