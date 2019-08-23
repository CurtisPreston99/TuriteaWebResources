package actions

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
)

func AddArticle(w http.ResponseWriter, r *http.Request) {
	p, id := se.checkPermission(r)
	switch p {
	case super:
	case normal:
		num, err := strconv.ParseInt(r.URL.Query().Get("num"), 16, 64)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = r.ParseForm()
		if err != nil {
			w.WriteHeader(400)
			return
		}
		data := r.Form.Get("data")
		articles, err := base.JsonToArticles(strings.NewReader(data), uint16(num))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		state := true
		for i, v := range articles {
			v.Id = base.GenArticleId()
			if !buffer.MainCache.CreateArticle(v) {
				_, _ = fmt.Fprint(w, i, " ")
				state = false
			}
		}
		if !state {
			_, _ = w.Write([]byte("-1"))
		}
		makeCookie(w, id)
		se.renew(id)
	case public:
		w.WriteHeader(401)
		// fixme do nothing or ?
	}
}
