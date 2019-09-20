package base

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int64
	Name string
	Role int
}

var userIdChan = make(chan int64, 100)
var userIdRecycle = make(chan int64, 100)

func userIdProvider() {
	var id int64 = 2
	for {
		select {
		case i := <-userIdRecycle:
			userIdChan <- i
		case userIdChan <- id:
			id++
		}
	}
}

func init() {
	go userIdProvider()
}

func GenUserId() int64 {
	return <-userIdChan
}

func RandomPassword() string {
	return fmt.Sprintf("%x", rand.Uint64())
}
func RecycleUserId(uid int64) {
	userIdRecycle <- uid
}
