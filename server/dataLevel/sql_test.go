package dataLevel

import (
	"testing"
)

/*
the other tests of the sql is making by hand
*/

func TestConnect(t *testing.T) {
	err := SQLWorker.Connect("postgres", "Turitea", "localhost", "turiteaSuper", "masseysuper")
	if err != nil {
		panic(err)
	}
}
