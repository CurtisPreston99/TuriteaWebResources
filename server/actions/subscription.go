package actions

import (
	"fmt"
	"log"
	"net/http"

	"TuriteaWebResources/server/dataLevel"
)

var speedControl = make(chan struct{}, 30)

func init() {
	for i := 0; i < 30; i++ {
		speedControl <- struct{}{}
	}
}

func addSubscription(w http.ResponseWriter, r *http.Request) {
	log.Println("call add subscription")
	<-speedControl
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	if len(name) | len(email) == 0 {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	if dataLevel.SQLWorker.CreatSubscription(name, email){
		_, _ = fmt.Fprint(w, "ok")
	} else {
		_, _ = fmt.Fprintf(w, "fail")
	}
	speedControl <-struct {}{}
}

func changeSubscription(w http.ResponseWriter, r *http.Request) {
	log.Println("call change subscription")
	<-speedControl
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	old := r.Form.Get("old")
	newOne := r.Form.Get("new")
	if len(old) | len(newOne) == 0 {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	if dataLevel.SQLWorker.ChangeSubscriptionEmail(old, newOne) {
		_, _ = fmt.Fprint(w, "ok")
	} else {
		_, _ = fmt.Fprintf(w, "fail")
	}
	speedControl <-struct {}{}
}

func deleteSubscription(w http.ResponseWriter, r *http.Request) {
	log.Println("call delete subscription")
	<-speedControl
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	email := r.Form.Get("email")
	if len(email) == 0 {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	if dataLevel.SQLWorker.DeleteSubscription(email) {
		_, _ = fmt.Fprint(w, "ok")
	} else {
		_, _ = fmt.Fprintf(w, "fail")
	}
	speedControl <-struct {}{}
}