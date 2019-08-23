package actions

import (
	"net/http"

	"TuriteaWebResources/server/dataLevel"
)

func AddFeedback(w http.ResponseWriter, r *http.Request) {
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
	dataLevel.SQLWorker.CreateFeedback(name, email, feedback)
	speedControl <-struct {}{}
}
