package actions

import (
	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/dataLevel"
	"fmt"
	"net/http"
)

func GetAllPins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call")
	err := base.PinsToJson(dataLevel.SQLNormal.GetAllPins(), w)
	if err != nil {
		fmt.Println(err)
	}
}
