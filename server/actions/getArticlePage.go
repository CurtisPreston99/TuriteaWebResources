package actions

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

var template []byte
func init() {
	f, err := os.Open("resources/html template/articleTemplate.html")
	if err != nil {
		panic(err)
	}
	s, err := f.Stat()
	if err != nil {
		panic(err)
	}
	template = make([]byte, s.Size())
	_, err = f.Read(template)
	if err != nil {
		panic(err)
	}
	f.Close()
}

func getArticlePage(w http.ResponseWriter, r *http.Request) {
	log.Println("get article")
	//todo send the template
	//hw := bufio.NewWriter(w)
	articleId := strings.SplitN(r.URL.Path, "/", 3)[2]
	if len(articleId) == 0 {
		http.Redirect(w, r, "../html/404.html", http.StatusTemporaryRedirect)
		return
	}
	id, err := strconv.ParseInt(articleId, 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	b, ok := buffer.MainCache.Load(base.ArticleKey(id))
	if !ok {
		http.Redirect(w, r, "../html/404.html", http.StatusTemporaryRedirect)
		return
	}
	a := b.(*base.Article)
	home := a.HomeContent
	http.SetCookie(w, &http.Cookie{Path:"/", Name:"home", Value:strconv.FormatInt(home, 16)})
	buffer.MainCache.LoadAsynchronous(dataLevel.ArticleContentKey(home))
	// for test
	//f, err := os.Open("resources/html template/articleTemplate.html")
	//if err != nil {
	//	panic(err)
	//}
	//
	//_ , err = io.Copy(w, f)
	//if err != nil {
	//	w.WriteHeader(500)
	//}
	//f.Close()

	//in deploy
	_, err = w.Write([]byte(fmt.Sprintf(string(template), a.Summary, a.Id)))
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
