package main

import (
	"fmt"
	"net/http"

	"TuriteaWebResources/server/actions"
)

//var fileHandle = http.FileServer(http.Dir("."))

func main() {
	config := &actions.Config{true, true}
	actions.Start(config)
	err := http.ListenAndServe("localhost:80", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
