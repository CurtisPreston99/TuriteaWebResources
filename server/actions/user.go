package actions

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"TuriteaWebResources/server/dataLevel"
)

func addUser(w http.ResponseWriter, r *http.Request) {
	log.Println("call add user")
	<-speedControl
	p, id := se.checkPermission(r)
	if p != super {
		w.WriteHeader(403)
		return
	} else {
		se.renew(id)
		makeCookie(w, id)
	}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	role, err := strconv.ParseInt(r.Form.Get("role"), 16, 64)
	name := r.Form.Get("name")
	if role > 2 || role < 1 || len(name) == 0 {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	password := dataLevel.SQLWorker.CreateRole(int(role), name)
	if len(password) == 0 {
		w.WriteHeader(500)
		speedControl <-struct {}{}
		return
	}
	_, _ = fmt.Fprintf(w, `{"name":"%s", "role":%d, "password":"%s"}`, name, role, password)
	speedControl <-struct {}{}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("call delete user")
	<-speedControl
	p, id := se.checkPermission(r)
	if p != super {
		w.WriteHeader(403)
		speedControl <-struct {}{}
		return
	} else {
		se.renew(id)
		makeCookie(w, id)
	}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	name := r.Form.Get("name")
	if len(name) == 0 {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	err = dataLevel.SQLWorker.DeleteUser(name)
	if err != nil {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	_, _ = fmt.Fprint(w, "ok")
	speedControl <-struct {}{}
}
