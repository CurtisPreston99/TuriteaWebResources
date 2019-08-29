package actions

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func init() {
	rand.Seed(1243494324)
}

func TestCreate(t *testing.T) {
	id, b := genToken(15)
	fmt.Println(b)
	b2, ok := parseBase(b)
	if !ok {
		t.Fatal()
	}
	if b2 != 2 {
		t.Fatal(b2)
	}
	id2, err := strconv.ParseInt(id, int(b2), 64)
	if err != nil {
		t.Fatal(err)
	}
	if id2 != 15 {
		t.Fatal()
	}
}
