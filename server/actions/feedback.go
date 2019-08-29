package actions

import (
	"log"
	"net/http"

	"TuriteaWebResources/server/dataLevel"
)

func AddFeedback(w http.ResponseWriter, r *http.Request) {
	log.Println("call add feedback")
	<-speedControl
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	email := r.Form.Get("email")
	name := r.Form.Get("name")
	feedback := r.Form.Get("feedback")
	if len(email) | len(name) | len(feedback) == 0 {
		w.WriteHeader(400)
		speedControl <-struct {}{}
		return
	}
	if !dataLevel.SQLWorker.CreateFeedback(name, email, feedback) {
		w.WriteHeader(500)
	}
	_, _ = w.Write([]byte("ok"))
	speedControl <-struct {}{}
}
