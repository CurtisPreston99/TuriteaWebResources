package actions

import (
	"math/rand"
	"net/http"
	"time"
)

func Start(config *Config) {
	rand.Seed(time.Now().Unix())

	ConfigAction(config)
	http.HandleFunc("/article/", getArticlePage)
	http.Handle("/", http.FileServer(http.Dir("static")))          //tested
	http.HandleFunc("/api/getPins", getPins)                       // tested
	http.HandleFunc("/api/getPin", getPin)                         // tested
	http.HandleFunc("/api/login", login)                           // tested
	http.HandleFunc("/api/addPins", addPins)                       // tested
	http.HandleFunc("/api/addArticle", addArticle)                 // tested
	http.HandleFunc("/api/addArticleFragment", addArticleFragment) // tested
	http.HandleFunc("/api/addImage", addImage)                     // tested
	http.HandleFunc("/api/update", update)
	http.HandleFunc("/api/delete", delete)                         // tested
	http.HandleFunc("/api/media", getMedia)                        // tested
	http.HandleFunc("/api/fragment", getFragment)                  // tested
	http.HandleFunc("/api/addSubscription", addSubscription)
	http.HandleFunc("/api/changeSubscription", changeSubscription)
	http.HandleFunc("/api/deleteSubscription", deleteSubscription)
	http.HandleFunc("/api/sendfeedback", addFeedback)           // tested
	http.HandleFunc("/api/addUser", addUser)                    // tested
	http.HandleFunc("/api/deleteUser", deleteUser)              // tested
	http.HandleFunc("/api/lastArticle", lastArticle)            // tested
	http.HandleFunc("/api/pinsByArticle", pinsByArticle)        // tested
	http.HandleFunc("/api/articlesByPin", articlesByPin)        // tested
	http.HandleFunc("/api/linkPinToArticle", linkArticleAndPin) // tested
	http.HandleFunc("/api/changePassword", changePassword)
	http.HandleFunc("/api/getImage", getImageLocal)
}
