package actions

import (
	"fmt"
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1243494327)  // it only a random number as key
}

func TestCreate(t *testing.T) {
	id, b := genToken(15)
	fmt.Println(b)
	b2, ok := parseBase(b)
	if !ok {
		t.Fatal()
	}
	fmt.Println(b2)
	id2, ok := parseToken(id, b)
	if !ok {
		t.Fatal()
	}
	if id2 != 15 {
		t.Fatal()
	}
}
