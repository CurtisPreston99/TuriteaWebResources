package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"TuriteaWebResources/server/actions"
)

//var fileHandle = http.FileServer(http.Dir("."))

func main() {
	config := &actions.Config{true, true}
	actions.Start(config)
	go func() {
		time.Tick(1*time.Second)
		(&http.Client{}).Get("http://localhost/api/getPins?north=-40&south=-41&east=176&west=175&timeBegin=0&timeEnd=20000")
		time.Tick(1*time.Second)
		log.Println("server start")
	}()
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
