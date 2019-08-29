package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"TuriteaWebResources/server/actions"
)

var fileHandle = http.FileServer(http.Dir("."))

func main() {
	rand.Seed(time.Now().Unix())
	config := actions.Config{true, true}
	actions.ConfigAction(config)
	http.HandleFunc("/article/", actions.GetArticlePage)
	http.Handle("/", http.FileServer(http.Dir("static"))) //tested
	http.HandleFunc("/api/getPins", actions.GetPins)  // tested
	http.HandleFunc("/api/getPin", actions.GetPin) // tested
	http.HandleFunc("/api/login", actions.Login)  // tested
	http.HandleFunc("/api/addPins", actions.AddPins)  // tested
	http.HandleFunc("/api/addArticle", actions.AddArticle)  // tested
	http.HandleFunc("/api/addArticleFragment", actions.AddArticleFragment) // tested
	http.HandleFunc("/api/addImage", actions.AddImage)  // tested
	http.HandleFunc("/api/update", actions.Update)
	http.HandleFunc("/api/delete", actions.Delete)  // tested
	http.HandleFunc("/api/media", actions.GetMedia) //tested
	http.HandleFunc("/api/fragment", actions.GetFragment)  // test
	http.HandleFunc("/api/addSubscription", actions.AddSubscription)
	http.HandleFunc("/api/changeSubscription",actions.ChangeSubscription)
	http.HandleFunc("/api/deleteSubscription", actions.DeleteSubscription)
	http.HandleFunc("/api/sendfeedback", actions.AddFeedback)  // tested
	http.HandleFunc("/api/addUser", actions.AddUser)  // tested
	http.HandleFunc("/api/deleteUser", actions.DeleteUser)  // tested
	http.HandleFunc("/api/lastArticle", actions.LastArticle)  // tested
	http.HandleFunc("/api/pinsByArticle", actions.PinsByArticle)  // tested
	http.HandleFunc("/api/articlesByPin", actions.ArticlesByPin)  // tested
	http.HandleFunc("/api/linkPinToArticle", actions.LinkArticleAndPin)  // tested
	// fixme the second parameter is set nil temporary change it later
	fmt.Println("server start")
	err := http.ListenAndServe("localhost:80", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
