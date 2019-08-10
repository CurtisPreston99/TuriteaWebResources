package dataLevel

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)


var r = &ArticleResource{Id:1, content: []byte("abc"), resourcesId: []int64{0}}

func TestSaveArticleContent(t *testing.T) {
	s := SaveArticleContentAndNotify(r)
	err := s()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadArticleContent(t *testing.T) {
	f := LoadArticleContent(ArticleKey(1))
	rt, err := f()
	if err != nil {
		t.Fatal()
	}
	if rs, ok := rt.(*ArticleResource); !ok {
		t.Fatal()
	} else {
		if rs.Id != r.Id {
			t.Fatal()
		}
		if string(rs.content) == string(r.content) {
			t.Fatal()
		}
		for i, v := range rs.resourcesId {
			if r.resourcesId[i] != v {
				t.Fatal()
			}
		}
	}
}

func init() {
	f, err := os.Open("test.jpg")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	ti.data = data
	d = data
}
var d []byte
var ti = &ImageResource{Id:1}

func TestLoadImage(t *testing.T) {
	f := LoadImage(ImageKey(1))
	i, err := f()
	if err != nil {
		t.Fatal(err)
	}
	if i, ok := i.(*ImageResource); ok {
		if i.Id == 1 && string(i.data) == string(d){
			fmt.Println(string(i.data))
			fmt.Println(string(d))
			return
		}
	}
	t.Fatal()
}

func TestSaveImage(t *testing.T) {
	s := SaveImageAndNotify(ti)
	err := s()
	if err != nil {
		t.Fatal()
	}
}