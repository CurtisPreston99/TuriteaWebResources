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

	//for test
	http.Handle("/", http.FileServer(http.Dir("../static")))       // auto tested

	// http.Handle("/", http.FileServer(http.Dir("static")))          // auto tested
	http.HandleFunc("/api/getPins", getPins)                       // auto tested
	http.HandleFunc("/api/getPin", getPin)                         // auto tested
	http.HandleFunc("/api/login", login)                           // auto tested
	http.HandleFunc("/api/addPins", addPins)                       // auto tested
	http.HandleFunc("/api/addArticle", addArticle)                 // auto tested
	http.HandleFunc("/api/addArticleFragment", addArticleFragment) // auto tested
	http.HandleFunc("/api/addImage", addImage)                     // auto tested
	http.HandleFunc("/api/update", update)                         // auto tested
	http.HandleFunc("/api/delete", deleteData)                     // tested
	http.HandleFunc("/api/media", getMedia)                        // auto tested
	http.HandleFunc("/api/fragment", getFragment)                  // auto tested
	http.HandleFunc("/api/addSubscription", addSubscription)
	http.HandleFunc("/api/changeSubscription", changeSubscription)
	http.HandleFunc("/api/deleteSubscription", deleteSubscription)
	http.HandleFunc("/api/sendfeedback", addFeedback)           // auto tested
	http.HandleFunc("/api/addUser", addUser)                    // auto tested
	http.HandleFunc("/api/deleteUser", deleteUser)              // auto tested
	http.HandleFunc("/api/lastArticle", lastArticle)            // auto tested
	http.HandleFunc("/api/pinsByArticle", pinsByArticle)        // auto tested
	http.HandleFunc("/api/articlesByPin", articlesByPin)        // auto tested
	http.HandleFunc("/api/linkPinToArticle", linkArticleAndPin) // auto tested
	http.HandleFunc("/api/unlinkArticleAndPin", unLinkPinAndArticle)
	http.HandleFunc("/api/changePassword", changePassword)      // auto tested
	http.HandleFunc("/api/getImage", getImageLocal)             // auto tested
}
