package actions

import (
	"fmt"
	"net/http"
)

func GetAllPins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call")
	//err := base.PinsToJson(dataLevel.SQLNormal.GetAllPins(), w)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
