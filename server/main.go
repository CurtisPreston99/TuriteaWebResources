package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var fileHandle = http.FileServer(http.Dir("."))

func main() {
	rand.Seed(time.Now().Unix())
	//config := actions.Config{true, true}
	//actions.ConfigAction(config)
	//http.HandleFunc("/api/getPins", actions.GetPins)
	//http.HandleFunc("/api/getPin", actions.GetPin)
	//http.Handle("/ct.html", fileHandle)
	//http.Handle("/cesiumTest2.html", fileHandle)
	//http.Handle("/cesiumtest2.js", fileHandle)
	//http.Handle("/ct2.js", fileHandle)
	//http.Handle("/cesium/", fileHandle)
	//http.Handle("./atricles/", http.FileServer(http.Dir(".")))
	//http.Handle("./resources/temPictures/", http.FileServer(http.Dir(".")))
	// fixme the second parameter is set nil temporary change it later
	err := http.ListenAndServe("localhost:80", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
