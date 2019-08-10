package dataLevel

import (
	"fmt"
	"testing"
)

/*
the other tests of the sql is making by hand
 */

func TestConnect(t *testing.T) {
	err := SQLSuper.Connect("postgres", "Turitea", "localhost", "turiteaSuper", "masseysuper")
	if err != nil {
		panic(err)
	}
	//m := base.Media{1, "a", "b", 1}
	_, err = SQLSuper.db.Query("insert into media (uid, title, url, type) VALUES (1, '11', 'http', 1)")
	if err != nil {
		fmt.Println(err)
		t.Fatal()
	}
}
