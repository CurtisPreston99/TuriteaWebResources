package dataLevel

import (
	"TuriteaWebResources/server/base"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

/*
	these query in these file had been test
 */

const (
	login = iota
	createRole
	deleteRole
	createSubscription
	deleteSubscription
	updateSubscriptionEmail
	createFeedback
	checkFeedback
	createPin
	deletePin
	createArticle
	loadArticle
	selectArticleByPin
	selectTopArticles
	selectNextTopArticles
	deleteArticle
	addMedia
	deleteMedia
	changeMediaUrl
	getMedia
	linkPinToArticle
	unlinkPinToArticle
	searchPinWithArticle
	getAllPins
	stmtLength
)

const (
	Public = iota
	Normal
	Super
)

var SQLNormal = &SqlLinker{}
var SQLPublic = &SqlLinker{}
var SQLSuper = &SqlLinker{}

func init() {
	err := SQLNormal.Connect("postgres", "Turitea", "localhost", "turiteaNormal", "massey")
	if err != nil {
		panic(err)
	}
	// todo add other roles
}

type SqlLinker struct {
	db *sql.DB
	stmtMap [stmtLength]*sql.Stmt
}

func (s *SqlLinker) Connect(driverName, dbName, host, userName, password string) (err error) {
	s.db, err = sql.Open(driverName, fmt.Sprintf("dbname=%s user=%s host=%s password=%s sslmode=disable", dbName, userName, host, password))
	if err != nil {
		return err
	}
	s.stmtMap[login], err = s.db.Prepare("select uid, role from users where name = $1 and password_hash = $2")
	if err != nil {
		return err
	}
	s.stmtMap[createRole], err = s.db.Prepare("insert into users (uid, name, password_hash, role) values ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	s.stmtMap[deleteRole], err = s.db.Prepare("delete from users where uid = $1")
	s.stmtMap[createSubscription], err = s.db.Prepare("insert into subscription (name, email) values ($1, $2)")
	if err != nil {
		return err
	}
	s.stmtMap[deleteSubscription], err = s.db.Prepare("delete from subscription where email = $1")
	if err != nil {
		return err
	}
	s.stmtMap[updateSubscriptionEmail], err = s.db.Prepare("update subscription set email = $1 where email = $2;")
	if err != nil {
		return err
	}
	s.stmtMap[createFeedback], err = s.db.Prepare("insert into feedback (id, name, feedback, email) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	s.stmtMap[checkFeedback], err = s.db.Prepare("update feedback set state = true where id = $1")
	if err != nil {
		return err
	}
	s.stmtMap[createPin], err = s.db.Prepare("insert into pins (uid, owner, latitude, longitude, time, tag_type, description) values ($1, $2, $3, $4, $5, $6, $7);")
	if err != nil {
		return err
	}
	s.stmtMap[getAllPins], err = s.db.Prepare("select uid, owner, latitude, longitude, time, tag_type, description from pins")
	if err != nil {
		return err
	}
	s.stmtMap[deletePin], err = s.db.Prepare("delete from pins where uid = $1")
	if err != nil {
		return err
	}
	s.stmtMap[createArticle], err = s.db.Prepare("insert into articles (id, summary, writenby) values ($1, $2, $3)")
	if err != nil {
		return err
	}
	s.stmtMap[loadArticle], err = s.db.Prepare("select summary, writenby from articles where id = $1")
	if err != nil {
		return err
	}
	// todo when add buffer change this one
	s.stmtMap[selectArticleByPin], err = s.db.Prepare("select id, summary, writenby from articles where id = (select article_id from pinlinkarticle where pin_id = $1)")
	if err != nil {
		return err
	}
	s.stmtMap[deleteArticle], err = s.db.Prepare("delete from articles where id=$1")
	if err != nil {
		return err
	}
	//fmt.Println("this one")
	s.stmtMap[selectTopArticles], err = s.db.Prepare("select id, summary, writenby from articles order by id desc limit $1")
	if err != nil {
		return err
	}
	// todo when add buffer change this one
	s.stmtMap[selectNextTopArticles], err = s.db.Prepare("select e.id, e.summary, e.writenby from (select id, summary, writenby from articles order by id desc limit $1) e order by id asc limit $2")
	if err != nil {
		return err
	}
	s.stmtMap[addMedia], err = s.db.Prepare("insert into media (uid, title, url, type) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	s.stmtMap[deleteMedia], err = s.db.Prepare("delete from media where uid = $1")
	if err != nil {
		return err
	}
	s.stmtMap[changeMediaUrl], err = s.db.Prepare("update media set url = $1 where uid = $2")
	if err != nil {
		return err
	}
	s.stmtMap[getMedia], err = s.db.Prepare("select url, title, type from media where uid = $1")
	if err != nil {
		return err
	}
	s.stmtMap[linkPinToArticle], err = s.db.Prepare("insert into pinlinkarticle (pin_id, article_id) values ($1, $2)")
	if err != nil {
		return err
	}
	s.stmtMap[unlinkPinToArticle], err = s.db.Prepare("delete from pinlinkarticle where pin_id = $1 and article_id = $2")
	if err != nil {
		return err
	}
	// todo when add buffer change this one
	s.stmtMap[searchPinWithArticle], err = s.db.Prepare("select uid, owner, latitude, longitude, time, tag_type from pins where uid = (select pin_id from pinlinkarticle where article_id = $1)")
	if err != nil {
		return err
	}
	return err
}

func (s *SqlLinker) Login(name string, password string) (user *base.User, ok bool) {
	rs, err := s.stmtMap[login].Query(name, password)
	if err != nil {
		return nil, false
	}
	if rs.Next() {
		var id int64
		var role int
		err = rs.Scan(&id, &role)
		if err != nil {
			err = rs.Close()
			return nil, false
		}
		user = base.GenUser(id, name, role, false)
	} else {
		err = rs.Close()
		return nil, false
	}
	if rs.Next() {
		log.Printf("sql server was attacked at %s with name=%s and password=%s", time.Now().Format("2016-01-02 15:04:05"), name, password)
	}
	err = rs.Close()
	return user, true
}

func (s *SqlLinker) CreateRole(role int, name string) error {
	// todo finish it
	user := base.GenUser(0, name, role, true)
	r, err := s.stmtMap[createRole].Query(user.Id, user.Name, base.RandomPassword(), role)
	if err != nil {
		base.RecycleUser(user, true)
		//pqErr := err.(*pq.Error)
		//pqErr.Code
	}
	err = r.Close()
	return nil
}

func (s *SqlLinker) DeleteUser(user *base.User) error {
	_, err := s.stmtMap[deleteRole].Query(user.Id)
	return err
}

func (s *SqlLinker) CreateSubscription(name, email string) bool {
	// todo finish it
	return false
}

func (s *SqlLinker) CreatePin(owner *base.User, latitude, longitude float64, time int64, tagType uint8, description string) (pin *base.Pin) {
	pin = base.GenPin(0, owner.Id, latitude, longitude, time, tagType, description, true)
	r, err :=s.stmtMap[createPin].Query(pin.Uid,owner.Id, latitude, longitude, time, tagType, description)
	if err != nil {
		base.RecyclePin(pin, true)
		err = r.Close()
		return nil
	}
	err = r.Close()
	return pin
}

func (s *SqlLinker) CreateArticle(summary string, writeBy int64) (article *base.Article) {
	article = base.GenArticle(0, writeBy, summary, true)
	r, err := s.stmtMap[createArticle].Query(article.Id, summary, writeBy)
	if err != nil {
		base.RecycleArticle(article, true)
		err = r.Close()
		return nil
	}
	err = r.Close()
	return article
}

func (s *SqlLinker) LoadArticle(id int64) *base.Article {
	r, err := s.stmtMap[loadArticle].Query(id)
	if err != nil {
		return nil
	}
	var summary string
	var writeBy int64
	err = r.Scan(&summary, &writeBy)
	if err != nil {
		err = r.Close()
		return nil
	}
	err = r.Close()
	return base.GenArticle(id, writeBy, summary, false)
}

func (s *SqlLinker) SelectArticlesWithPin(pin *base.Pin) []*base.Article {
	r, err := s.stmtMap[selectArticleByPin].Query(pin.Uid)
	if err != nil {
		return nil
	}
	var summary string
	var writeBy int64
	var id int64
	var goal = make([]*base.Article, 0, 2)
	for r.Next() {
		err = r.Scan(&id, &writeBy, &summary)
		if err != nil {
			err = r.Close()
			return nil
		}
		goal = append(goal, base.GenArticle(id, writeBy, summary, false))
	}
	err = r.Close()
	return goal
}

func (s *SqlLinker) SelectTopArticles(top uint8) []*base.Article {
	r, err := s.stmtMap[selectTopArticles].Query(top)
	if err != nil {
		return nil
	}
	var id, writeBy int64
	var summary string
	var goal = make([]*base.Article, 0, top)
	for r.Next(){
		err = r.Scan(&id, &summary, &writeBy)
		if err != nil {
			err = r.Close()
			return nil
		}
		goal = append(goal, base.GenArticle(id, writeBy, summary, false))
	}
	err = r.Close()
	return goal
}

func (s *SqlLinker) SelectNextTopArticles(begin int64, length uint8) []*base.Article {
	if length == 0 || begin < 0 {
		return []*base.Article{}
	}
	r, err := s.stmtMap[selectNextTopArticles].Query(begin+int64(length), length)
	if err != nil {
		return nil
	}
	var id, writeBy int64
	var summary string
	goal := make([]*base.Article, 0, length)
	if r.Next() {
		err = r.Scan(&id, &summary, &writeBy)
		if err != nil {
			err = r.Close()
			return nil
		}
		goal = append(goal, base.GenArticle(id, writeBy, summary,false))
	}
	err = r.Close()
	return goal
}

func (s *SqlLinker) GetMedia(id int64) (media *base.Media) {
	r, err := s.stmtMap[getMedia].Query(id)
	if err != nil {
		return nil
	}
	if r.Next() {
		var t uint8
		var title string
		var url string
		err = r.Scan(&url, &title, &t)
		if err != nil {
			err = r.Close()
			return nil
		} else {
			media = base.GenMedia(id, t, title, url, false)
		}
	} else {
		err = r.Close()
		return nil
	}
	if r.Next() {
		log.Printf("sql server was attacked at %s with id=%d", time.Now().Format("2016-01-02 15:04:05"), id)
	}
	err = r.Close()
	return media
}

func (s *SqlLinker) AddMedia(title, url string, t uint8) (media *base.Media) {
	media = base.GenMedia(0, t, title, url, true)
	r, err := s.stmtMap[addMedia].Query(media.Uid, title, url, t)
	if err != nil {
		err = r.Close()
		base.RecycleMedia(media, true)
		return nil
	}
	err = r.Close()
	return media
}

func (s *SqlLinker) LinkPinToArticle(pin *base.Pin, article *base.Article) bool {
	r, err := s.stmtMap[linkPinToArticle].Query(pin.Uid, article.Id)
	if err != nil {
		return false
	}
	_ = r.Close()
	return true
}

func (s *SqlLinker) UnLinkPinToArticle(pin *base.Pin, article *base.Article) bool {
	r, err := s.stmtMap[unlinkPinToArticle].Query(pin.Uid, article.Id)
	if err != nil {
		return false
	}
	_ = r.Close()
	return true
}

func (s *SqlLinker) SearchPinsWithArticle(article *base.Article) []*base.Pin {
	r, err := s.stmtMap[searchPinWithArticle].Query(article.Id)
	if err != nil {
		return nil
	}
	var owner, id, t int64
	var lat, lon float64
	var tagType uint8
	var description string
	goal := make([]*base.Pin, 0, 2)
	for r.Next() {
		err = r.Scan(&id, &owner, &lat, &lon, &t, &tagType, &description)
		if err != nil {
			return nil
		}
		goal = append(goal, base.GenPin(id, owner, lat, lon, t, tagType, description, false))
	}
	_ = r.Close()
	return goal
}

// fixme it just for test
func (s *SqlLinker) GetAllPins() []*base.Pin {
	r, err := s.stmtMap[getAllPins].Query()
	if err != nil {
		_ = r.Close()
		return nil
	}
	var uid, owner, t int64
	var lat, long float64
	var tagType uint8
	var description string
	goal := make([]*base.Pin,0, 20)
	for r.Next() {
		err = r.Scan(&uid, &owner, &lat, &long, &t, &tagType, &description)
		if err != nil {
			return nil
		}
		goal = append(goal, base.GenPin(uid, owner, lat, long, t, tagType, description, false))
	}
	return goal
}