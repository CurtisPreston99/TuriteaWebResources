package base

import (
	"fmt"
	"testing"
)

/*
	because it is no different for these Gen** and Recycle** functions so just test one
 */

func TestCreate(t *testing.T) {
	a := GenArticle(1, 1, 1, "11")
	a.Id = GenArticleId()
	if a.Id != 2 || a.WriteBy != 1 || a.Summary != "11" {
		fmt.Printf("%v\n", a)
		t.Fatal()
	}
}

func TestGen(t *testing.T) {
	a := GenArticle(4, 1, 1, "11")
	if a.Id != 4 || a.WriteBy != 1 || a.Summary != "11" {
		fmt.Printf("%v\n", a)
		t.Fatal()
	}
}

func TestRecycleDelete(t *testing.T) {
	a := &Article{1, 1, "11", 1}
	RecycleArticle(a, true)
	for i := 0; i < 200; i++ {
		if id := <- articleIdChan; id == 1 {
			return
		}
	}
	t.Fatal()
}

func TestRecycle(t *testing.T) {
	a := &Article{0, 1, "11", 1}
	RecycleArticle(a, false)
	for i := 0; i < 200; i++ {
		if id := <- articleIdChan; id == 0 {
			t.Fatal()
		}
	}
}
