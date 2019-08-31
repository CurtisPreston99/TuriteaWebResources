package actions

import (
	"net/http"
)

func ControlSetting(w http.ResponseWriter, r *http.Request) {
	p, id := se.checkPermission(r)
	if p == public {
		// todo maybe it will redirect
		w.WriteHeader(401)
		return
	} else {
		// todo give the setting html
	}
	se.renew(id)
	makeCookie(w, id)
}
