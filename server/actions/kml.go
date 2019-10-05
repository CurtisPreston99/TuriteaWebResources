package actions

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)
var lock *sync.RWMutex
var kmlList map[string]uint8

func init() {
	loadAllKMLName()
	lock = new(sync.RWMutex)
}
var kmlPath = "./resources/kml/"
func loadAllKMLName() {
	files, err := ioutil.ReadDir(kmlPath[:len(kmlPath)-1])
	if err != nil {
		panic(err)
	}
	kmlList = map[string]uint8{}
	for _, v := range files {
		kmlList[v.Name()] = 0
	}
}

func putKml(w http.ResponseWriter, r *http.Request) {
	log.Println("call put kml")
	p, uid := se.checkPermission(r)
	switch p {
	case normal:
		fallthrough
	case super:
		f, head, err := r.FormFile("kml")
		if err != nil {
			w.WriteHeader(400)
			return
		}
		title := head.Filename
		k, err := os.Create(kmlPath + title)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		_, err = io.Copy(k, f)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		err = r.MultipartForm.RemoveAll()
		if err != nil {
			log.Println(time.Now().Format(time.Stamp), err)
		}
		err = f.Close()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		makeCookie(w, uid)
		se.renew(uid)
		lock.Lock()
		kmlList[title] = 0
		lock.Unlock()
		http.Redirect(w, r, "../html/settings.html#tabs-5", http.StatusTemporaryRedirect)
	case public:
		w.WriteHeader(401)
	}
}

func deleteKML(w http.ResponseWriter, r *http.Request) {
	log.Println("call put kml")
	p, uid := se.checkPermission(r)
	switch p {
	case normal:
		fallthrough
	case super:
		vs := r.URL.Query()
		title := vs.Get("name")
		err := os.Remove(kmlPath + title)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		makeCookie(w, uid)
		se.renew(uid)
		lock.Lock()
		delete(kmlList, title)
		lock.Unlock()
	case public:
		w.WriteHeader(401)
	}
}

func listKML(w http.ResponseWriter, r *http.Request) {
	log.Println("call list kml")
	_, err := w.Write([]byte("["))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	c := 0
	lock.RLock()
	defer lock.RUnlock()
	for k := range kmlList {
		_, err = w.Write([]byte("\""))
		if err != nil {
			w.WriteHeader(500)
			return
		}
		_, err = w.Write([]byte(k))
		if err != nil {
			w.WriteHeader(500)
			return
		}
		_, err = w.Write([]byte("\""))
		if err != nil {
			w.WriteHeader(500)
			return
		}
		c++
		if c != len(kmlList) {
			_, err = w.Write([]byte(","))
			if err != nil {
				w.WriteHeader(500)
				return
			}
		}
	}
	_, err = w.Write([]byte("]"))
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func getKml(w http.ResponseWriter, r *http.Request) {
	log.Println("call get kml")
	vs := r.URL.Query()
	title := vs.Get("name")
	lock.RLock()
	if _, ok := kmlList[title]; !ok {
		w.WriteHeader(404)
		lock.RUnlock()
		return
	}
	lock.RUnlock()
	w.Header().Set("Content-type", "text/xml")
	w.Header().Set("Cache-Control", "max-age=31536000")
	var f *os.File
	var err error
	count := 0
	retry:
	f, err = os.Open(kmlPath + title)
	if err != nil {
		if count <= 3{
			count++
			goto retry
		}
		w.WriteHeader(404)
		return
	}
	_, err = io.Copy(w, f)
	if err != nil {
		if count <= 3{
			count++
			goto retry
		}
		w.WriteHeader(500)
		return
	}
	_ = f.Close()
}
