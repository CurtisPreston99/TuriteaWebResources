package actions

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/dataLevel"
)

const (
	public = iota
	normal
	super
)

type sessions struct {
	lock *sync.RWMutex
	session map[int64]*session
}

func (s *sessions) sessionsCleaner() {
	ticker := time.NewTicker(17 * time.Minute)
	c := ticker.C
	for {
		<-c
		s.lock.Lock()
		for k, v := range s.session {
			if v.lastTime + d < time.Now().Unix() {
				delete(s.session, k)
			}
		}
		s.lock.Unlock()
	}
}

type session struct {
	user *base.User
	lastTime int64
}
const d = int64(5 * time.Minute / time.Second)
func (s *sessions) checkPermission(r *http.Request) (uint8, int64) {
	c, err := r.Cookie("lastTime")
	if err != nil {
		return public, -1
	}
	id := c.Value
	c, err = r.Cookie("key")
	b := c.Value
	i, ok := parseToken(id, b)
	if !ok {
		return public, -1
	}
	s.lock.RLock()
	one, ok := s.session[i]
	s.lock.RUnlock()
	if !ok {
		return public, -1
	}
	if one.lastTime + d < time.Now().Unix() {
		s.lock.Lock()
		delete(s.session, i)
		s.lock.Unlock()
		return public, -1
	}
	return uint8(one.user.Role), one.user.Id
}

func (s *sessions) renew(id int64) {
	s.lock.RLock()
	s.session[id].lastTime = time.Now().Unix()
	s.lock.RUnlock()
}

var se = &sessions{new(sync.RWMutex), make(map[int64]*session)}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("call login")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(404)
		return
	}
	name := r.Form.Get("name")
	pw := r.Form.Get("pw")
	u := dataLevel.SQLWorker.Login(name, pw)
	if u != nil {
		makeCookie(w, u.Id)
		se.lock.Lock()
		se.session[u.Id] = &session{u, time.Now().Unix()}
		se.lock.Unlock()
		if u.Role == super {
			http.SetCookie(w, &http.Cookie{Path:"/", Value:"true", Name:"super"})
			//http.Redirect(w, r, "../super/control.html", 307)
			w.WriteHeader(200)
		} else if u.Role == normal {
			//http.Redirect(w, r, "../normal/control.html", 307)
			w.WriteHeader(200)
		}
	} else {
		//http.Redirect(w, r, "../others/loginFail.html", 307)
		w.WriteHeader(401)
	}
}

func genToken(uid int64) (string, string) {
	r := rand.Int31n(33) + 3
	id := strconv.FormatInt(uid + 13, int(r))
	baseNumber := getBase(uint8(r))
	return id, baseNumber
}

func getBase(baseNumber uint8) string {
	goal := uint64(0)
	for i := 0; i < 8; i ++ {
		if baseNumber & 0x1 == 1 {
			r := rand.Int63n(256)
			//r := 0
			goal |= uint64(byte(r) | (1 << uint8(i)))
		} else {
			r := rand.Int63n(256)
			//r := 0
			goal |= uint64(byte(r) &^ (1 << uint8(i)))
		}
		if i < 7 {
			goal <<= 8
			baseNumber >>= 1
		}
	}
	return strconv.FormatInt(int64(goal), 15)
}

func parseToken(id, b string) (int64, bool) {
	baseNumber, ok := parseBase(b)
	if !ok {
		return 0, false
	}
	goal, err := strconv.ParseInt(id, int(baseNumber), 64)
	if err != nil {
		return 0, false
	}
	return goal - 13, true
}

func parseBase(o string) (uint8, bool) {
	num, err := strconv.ParseInt(o, 15, 64)
	if err != nil {
		return 0, false
	}
	var goal = uint8(0)
	for i := 0; i < 8; i++ {
		goal |= uint8(num) & (1 << uint8(7-i))
		num >>= 8
	}
	if goal < 2 || goal > 36 {
		return 0, false
	}
	return goal, true
}


func makeCookie(w http.ResponseWriter, uid int64) {
	id, key := genToken(uid)
	http.SetCookie(w, &http.Cookie{Path:"/", Name: "lastTime", Value: id, HttpOnly: true})
	http.SetCookie(w, &http.Cookie{Path:"/", Name: "key", Value: key, HttpOnly: true})
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	p, id := se.checkPermission(r)
	if p == public {
		w.WriteHeader(403)
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	old := r.Form.Get("old")
	n := r.Form.Get("new")
	if len(old) == 0 || len(n) == 0 {
		w.WriteHeader(400)
		return
	}
	ok := dataLevel.SQLWorker.ChangePassword(old, n, id)
	if !ok {
		w.WriteHeader(500)
		return
	}
	se.renew(id)
	makeCookie(w, id)
}
