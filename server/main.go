package main

import (
	"TuriteaWebResources/server/actions"
	"math/rand"
	"net/http"
	"time"
)

var fileHandle = http.FileServer(http.Dir("."))

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/api/pins", actions.GetAllPins)
	http.Handle("/ct.html", fileHandle)
	http.Handle("/cesiumTest2.html", fileHandle)
	http.Handle("/cesiumtest2.js", fileHandle)
	http.Handle("/ct2.js", fileHandle)
	http.Handle("/cesium/", fileHandle)
	//http.Handle("./atricles/", http.FileServer(http.Dir(".")))
	//http.Handle("./resources/temPictures/", http.FileServer(http.Dir(".")))
	// fixme the second parameter is set nil temporary change it later
	http.ListenAndServe("localhost:8080", nil)
}
