package base

import (
	"fmt"
	"math/rand"
	"sync"
)

type User struct {
	Id int64
	Name string
	Role int
}

var userPool = new(sync.Pool)
var userIdChan = make(chan int64, 100)
var userIdRecycle = make(chan int64, 100)

func userIdProvider() {
	var id int64
	for {
		select {
		case i := <-userIdRecycle:
			userIdChan <- i
		case userIdChan<-id:
			id ++
		}
	}
}

func init() {
	userPool.New = func() interface{} {
		return &User{}
	}
	go userIdProvider()
}


func GenUser(Id int64, name string, role int, newOne bool) *User {
	if newOne {
		Id = <-userIdChan
	}
	user := userPool.Get().(*User)
	user.Id = Id
	user.Name = name
	user.Role = role
	return user
}

func RandomPassword() string {
	return fmt.Sprintf("%x", rand.Uint64())
}

func RecycleUser(user *User, delete bool) {
	userPool.Put(user)
	if delete {
		userIdRecycle <- user.Id
	}
}