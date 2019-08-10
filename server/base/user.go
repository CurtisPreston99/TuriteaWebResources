package base

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id int64
	Name string
	Role int
}

func GenUser(Id int64, name string, role int, newOne bool) *User {
	// todo finish it
	return nil
}

func RandomPassword() string {
	return fmt.Sprintf("%x", rand.Uint64())
}

func RecycleUser(user *User, delete bool) {

}