package actions

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
)

func addPins(w http.ResponseWriter, r *http.Request) {
	log.Println("call add pin")
	p, id := se.checkPermission(r)
	switch p {
	case super:
		fallthrough
	case normal:
		num, err := strconv.ParseInt(r.URL.Query().Get("num"), 16, 64)
		if err != nil {
			return
		}
		err = r.ParseForm()
		if err != nil {
			return
		}
		data := r.Form.Get("data")
		pins, err := base.JsonToPins(strings.NewReader(data), uint16(num))
		if err != nil {
			w.WriteHeader(400)
			return
		}
		state := true
		for i, v := range pins {
			v.Uid = base.GenPinId()
			if !buffer.MainCache.CreatePin(v) {
				_, _ = fmt.Fprint(w, i, " ")
				base.RecyclePin(v, true)
				state = false
			}
		}
		if !state {
			_, _ = w.Write([]byte("-1"))
		} else {
			_, _ = w.Write([]byte("ok"))
		}
		makeCookie(w, id)
		se.renew(id)
	case public:
		w.WriteHeader(403)
		// fixme do nothing or ?
	}
}
