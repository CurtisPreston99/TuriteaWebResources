package dataLevel

import (
	"github.com/ChenXingyuChina/asynchronousIO"
	"TuriteaWebResources/server/base"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)


var r = &ArticleResource{Id:1, Content: []byte("abc"), ResourcesId: []Resource{{1,1}, {1,3}}}

func TestSaveArticleContent(t *testing.T) {
	s := SaveArticleContentAndNotify(r)
	err := s()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadArticleContent(t *testing.T) {
	f := LoadArticleContent(ArticleContentKey(1))
	rt, err := f()
	if err != nil {
		t.Fatal(err)
	}
	if rs, ok := rt.(*ArticleResource); !ok {
		t.Fatal()
	} else {
		if rs.Id != r.Id {
			t.Fatal()
		}
		if strings.Compare(string(rs.Content), string(r.Content)) != 0{
			t.Fatal(string(rs.Content))
		}
		for i, v := range rs.ResourcesId {
			if r.ResourcesId[i] != v {
				t.Fatal()
			}
		}
	}
}

func TestDeleteArticle(t *testing.T) {
	f := DeleteArticle(base.ArticleKey(1))
	err := f()
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	OnLoadResourceId = func(resources []Resource) {

	}
	OnLoadMedia = func(key ImageKey) {

	}
	f, err := os.Open("test.jpg")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	ti.Data = data
	d = data
	Init()
}
var d []byte

var ti = &ImageResource{Id:1}

func TestSaveImage(t *testing.T) {
	s := SaveImageAndNotify(ti)
	err := s()
	if err != nil {
		t.Fatal()
	}
}

func TestLoadImage(t *testing.T) {
	f := LoadImage(ImageKey(1))
	i, err := f()
	if err != nil {
		t.Fatal(err)
	}
	if i, ok := i.(*ImageResource); ok {
		if i.Id == 1 && string(i.Data) == string(d){
			//fmt.Println(string(i.Data))
			//fmt.Println(string(d))
			return
		}
	}
	t.Fatal()
}

func TestSqlLinker_CreateArticle(t *testing.T) {
	ok := SQLWorker.CreateArticle("abccc", 0, 0, 1)
	if !ok {
		t.Fatal()
	}
}

func TestLoadArticle(t *testing.T) {
	f := Load(base.ArticleKey(0))
	var err error
	b, err = f()
	if err != nil {
		fmt.Printf("%t\n", err)
		t.Fatal(err)
	}
	fmt.Println(b)
}

func TestSaveArticle(t *testing.T) {
	b.(*base.Article).Summary = "test"
	f := SaveAndNotify(b)
	err := f()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqlLinker_DeleteArticle(t *testing.T) {
	if err := SQLWorker.DeleteArticle(0); err != nil {
		t.Fatal(err)
	}
}

var b asynchronousIO.Bean
func TestLoadPin(t *testing.T) {
	f := LoadPin(base.PinKey(1))
	var err error
	b, err = f()
	if err != nil {
		fmt.Printf("%s", err)
		t.Fatal()
	}
	fmt.Println(b)
}

func TestSavePin(t *testing.T) {
	TestLoadPin(t)
	pin := b.(*base.Pin)
	pin.Description = "for test"
	f := SavePinAndNotify(pin)
	err := f()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqlLinker_CreatePin(t *testing.T) {
	if !SQLWorker.CreatePin(27, 0, 10, 10, 111, base.TagNameToNumber["hospital"], "test", "abc", "#00ff00") {
		t.Fatal()
	}
}

func TestDeletePin(t *testing.T) {
	f := DeletePin(base.PinKey(27))
	err := f()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqlLinker_AddMedia(t *testing.T) {
	if !SQLWorker.AddMedia(0, "2333", "23333", 0) {
		t.Fatal()
	}
}

func TestLoadMedia(t *testing.T) {
	f := LoadMedia(base.MediaKey(0))
	var err error
	b, err = f()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}

func TestSaveMedia(t *testing.T) {
	TestLoadMedia(t)
	media := b.(*base.Media)
	media.Title = "2233"
	f := SaveMediaAndNotify(media)
	err := f()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMedia(t *testing.T) {
	f := DeleteMedia(base.MediaKey(0))
	err := f()
	if err != nil {
		t.Fatal()
	}
}
