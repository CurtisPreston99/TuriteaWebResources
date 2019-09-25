package dataLevel

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"TuriteaWebResources/server/base"
)

/*
	these query in these file had been test
 */

const (
	login = iota
	createRole
	deleteRole
	allRole
	changePassword
	createSubscription
	deleteSubscription
	updateSubscriptionEmail

	createFeedback
	checkFeedback

	createPin
	deletePin
	getPinById
	updatePin
	getAllPins
	getPinsInArea

	createArticle
	loadArticle
	updateArticle
	selectArticleIdByPin
	selectTopArticles
	selectNextTopArticles
	deleteArticle

	addMedia
	deleteMedia
	changeMedia
	getMedia

	linkPinToArticle
	unlinkPinToArticle
	searchPinIdWithArticle

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
var SQLWorker = &SqlLinker{}

func init() {
	err := SQLNormal.Connect("postgres", "Turitea", "localhost", "turiteaNormal", "massey")
	if err != nil {
		panic(err)
	}
	err = SQLSuper.Connect("postgres", "Turitea", "localhost", "turiteaSuper", "masseysuper")
	if err != nil {
		panic(err)
	}
	err = SQLPublic.Connect("postgres", "Turitea", "localhost", "turiteaPublic", "masseyPublic")
	if err != nil {
		panic(err)
	}
	err = SQLWorker.Connect("postgres", "Turitea", "localhost", "turiteaWorker", "tutiteaworker")
	if err != nil {
		panic(err)
	}
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
	s.stmtMap[changePassword], err = s.db.Prepare("update users set password_hash = $1 where uid=$2 and password_hash=$3;")
	if err != nil {
		return err
	}
	s.stmtMap[deleteRole], err = s.db.Prepare("delete from users where name = $1")
	if err != nil {
		return err
	}
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
	s.stmtMap[createFeedback], err = s.db.Prepare("insert into feedback (name, feedback, email, state) VALUES ($1, $2, $3, false)")
	if err != nil {
		return err
	}
	s.stmtMap[checkFeedback], err = s.db.Prepare("update feedback set state = true where id = $1")
	if err != nil {
		return err
	}
	s.stmtMap[createPin], err = s.db.Prepare("insert into pins (uid, owner, latitude, longitude, time, tag_type, description, name, color) values ($1, $2, $3, $4, $5, $6, $7, $8, $9);")
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
	s.stmtMap[createArticle], err = s.db.Prepare("insert into articles (id, summary, writenby, home_content) values ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	s.stmtMap[loadArticle], err = s.db.Prepare("select summary, writenby, home_content from articles where id = $1")
	if err != nil {
		return err
	}
	s.stmtMap[selectArticleIdByPin], err = s.db.Prepare("select article_id from pinlinkarticle where pin_id = $1")
	if err != nil {
		return err
	}
	s.stmtMap[deleteArticle], err = s.db.Prepare("delete from articles where id=$1")
	if err != nil {
		return err
	}
	//fmt.Println("this one")
	s.stmtMap[selectTopArticles], err = s.db.Prepare("select id from articles order by id desc limit $1")
	if err != nil {
		return err
	}
	s.stmtMap[selectNextTopArticles], err = s.db.Prepare("select e.id from (select id from articles order by id desc limit $1) e order by id asc limit $2")
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
	s.stmtMap[changeMedia], err = s.db.Prepare("update media set url = $1, title = $2 where uid = $3")
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
	s.stmtMap[searchPinIdWithArticle], err = s.db.Prepare("select pin_id from pinlinkarticle where article_id = $1;")
	if err != nil {
		return err
	}
	s.stmtMap[getPinById], err = s.db.Prepare("select uid, owner, latitude, longitude, time, description, tag_type, name, color from pins where uid = $1")
	if err != nil {
		return err
	}
	s.stmtMap[updatePin], err = s.db.Prepare("update pins set tag_type = $1, name = $2, description = $3, color = $4 where uid = $5;")
	if err != nil {
		return err
	}
	s.stmtMap[updateArticle], err = s.db.Prepare("update articles set summary = $1 where id = $2;")
	if err != nil {
		return err
	}
	s.stmtMap[getPinsInArea], err = s.db.Prepare("select uid from pins where (latitude between $1 and $2) and (longitude between $3 and $4) and (time between $5 and $6)")
	if err != nil {
		return err
	}
	s.stmtMap[allRole], err = s.db.Prepare("select name, role from users")
	return err
}

func (s *SqlLinker) Login(name string, password string) *base.User {
	rs, err := s.stmtMap[login].Query(name, password)
	if err != nil {
		return nil
	}
	var id int64
	var role int
	if rs.Next() {
		err = rs.Scan(&id, &role)
		if err != nil {
			err = rs.Close()
			return nil
		}
	} else {
		err = rs.Close()
		return nil
	}
	if rs.Next() {
		log.Printf("sql server was attacked at %s with name=%s and password=%s", time.Now().Format("2016-01-02 15:04:05"), name, password)
	}
	err = rs.Close()
	return &base.User{id, name, role}
}

func (s *SqlLinker) CreateRole(role int, name string) string {
	userId := base.GenUserId()
	passWord := base.RandomPassword()
	//fmt.Println(fmt.Sprintf("%x", md5.New().Sum([]byte(passWord))))
	r, err := s.stmtMap[createRole].Query(userId, name, fmt.Sprintf("%x", md5.Sum([]byte(passWord))), role)
	if err != nil {
		base.RecycleUserId(userId)
		return ""
	}
	err = r.Close()
	return passWord
}

func (s *SqlLinker) DeleteUser(name string) error {
	_, err := s.stmtMap[deleteRole].Query(name)
	return err
}

func (s *SqlLinker) ChangePassword(passwordHash, newPassword string, uid int64) bool {
	_, err := s.stmtMap[changePassword].Query(newPassword, uid, passwordHash)
	if err != nil {
		return false
	}
	return true
}

func (s *SqlLinker) CreatePin(id, owner int64, latitude, longitude float64, t int64, tagType uint8, description, name, color string) bool{
	r, err :=s.stmtMap[createPin].Query(id, owner, latitude, longitude, t, tagType, description, name, color)
	if err != nil {
		err = r.Close()
		return false
	}
	err = r.Close()
	return true
}

func (s *SqlLinker) CreateArticle(summary string, id, writeBy, homeContent int64) bool {
	r, err := s.stmtMap[createArticle].Query(id, summary, writeBy, homeContent)
	if err != nil {
		return false
	}
	err = r.Close()
	return true
}

func (s *SqlLinker) LoadArticle(id int64) *base.Article {
	r, err := s.stmtMap[loadArticle].Query(id)
	if err != nil {
		return nil
	}
	var summary string
	var writeBy, home int64
	r.Next()
	err = r.Scan(&summary, &writeBy, &home)
	if err != nil {
		err = r.Close()
		return nil
	}
	err = r.Close()
	return base.GenArticle(id, writeBy, home, summary)
}

func (s *SqlLinker) SelectArticlesIdWithPin(pinId int64) []int64 {
	r, err := s.stmtMap[selectArticleIdByPin].Query(pinId)
	if err != nil {
		return nil
	}
	var id int64
	var goal = make([]int64, 0, 2)
	for r.Next() {
		err = r.Scan(&id)
		if err != nil {
			err = r.Close()
			return nil
		}
		goal = append(goal, id)
	}
	err = r.Close()
	return goal
}

func (s *SqlLinker) SelectTopArticles(top uint8) []int64 {
	r, err := s.stmtMap[selectTopArticles].Query(top)
	if err != nil {
		return nil
	}
	var id int64
	var goal = make([]int64, 0, top)
	for r.Next(){
		err = r.Scan(&id)
		if err != nil {
			err = r.Close()
			return nil
		}
		goal = append(goal, id)
	}
	err = r.Close()
	return goal
}

func (s *SqlLinker) SelectNextTopArticles(begin int64, length uint8) []int64 {
	if length == 0 || begin < 0 {
		return []int64{}
	}
	r, err := s.stmtMap[selectNextTopArticles].Query(begin+int64(length), length)
	if err != nil {
		return nil
	}
	var id int64
	goal := make([]int64, 0, length)
	if r.Next() {
		err = r.Scan(&id)
		if err != nil {
			err = r.Close()
			return nil
		}
		goal = append(goal, id)
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
			media = base.GenMedia(id, t, title, url)
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

func (s *SqlLinker) AddMedia(id int64, title, url string, t uint8) bool {
	r, err := s.stmtMap[addMedia].Query(id, title, url, t)
	if err != nil {
		return false
	}
	err = r.Close()
	return true
}

func (s *SqlLinker) LinkPinToArticle(pid, aid int64) bool {
	r, err := s.stmtMap[linkPinToArticle].Query(pid, aid)
	if err != nil {
		return false
	}
	_ = r.Close()
	return true
}

func (s *SqlLinker) UnLinkPinToArticle(pinId , articleId int64) bool {
	r, err := s.stmtMap[unlinkPinToArticle].Query(pinId, articleId)
	if err != nil {
		return false
	}
	_ = r.Close()
	return true
}

func (s *SqlLinker) SearchPinsIdWithArticle(article int64) []int64 {
	r, err := s.stmtMap[searchPinIdWithArticle].Query(article)
	if err != nil {
		return nil
	}
	var id int64
	goal := make([]int64, 0, 2)
	for r.Next() {
		err = r.Scan(&id)
		if err != nil {
			return nil
		}
		goal = append(goal, id)
	}
	_ = r.Close()
	return goal
}

func (s *SqlLinker) CreatSubscription(name, email string) bool {
	r, err := s.stmtMap[createSubscription].Query(name, email)
	_ = r.Close()
	if err != nil {
		return false
	}
	return true
}

func (s *SqlLinker) DeleteSubscription(email string) bool {
	r, err := s.stmtMap[deleteSubscription].Query(email)
	if err != nil {
		return false
	}
	_ = r.Close()
	return true
}

func (s *SqlLinker) ChangeSubscriptionEmail(emailOld, emailNew string) bool {
	_, err := s.stmtMap[updateSubscriptionEmail].Query(emailNew, emailOld)
	if err != nil {
		return false
	}
	return true
}

func (s *SqlLinker) CreateFeedback(name, email, feedback string) bool {
	_, err := s.stmtMap[createFeedback].Query(name, feedback, email)
	if err != nil {
		return false
	}
	return true
}

func (s *SqlLinker) CheckFeedback(id int64) bool {
	_, err := s.stmtMap[checkFeedback].Query(id)
	if err != nil {
		return false
	}
	return true
}

func (s *SqlLinker) GetPinById(id int64) (*base.Pin, bool) {
	var uid, owner, t int64
	var description, name, color string
	var lat, long float64
	var tagType uint8
	r, err := s.stmtMap[getPinById].Query(id)
	if err != nil {
		return nil, false
	}
	if r.Next() {
		err = r.Scan(&uid, &owner, &lat, &long, &t, &description, &tagType, &name, &color)
		if err != nil {
			return nil, false
		}
	}
	return base.GenPin(uid, owner, lat, long, t, tagType, description, name, color), true
}

func (s *SqlLinker) UpdatePin(pin *base.Pin) bool {
	_, err := s.stmtMap[updatePin].Query(base.TagNameToNumber[pin.TagType], pin.Name, pin.Description, pin.Color, pin.Uid)
	if err != nil {
		return false
	}
	return true
}

func (s *SqlLinker) DeletePin(uid int64) bool {
	_, err := s.stmtMap[deletePin].Query(uid)
	if err != nil {
		return false
	}
	return true
}

func (s *SqlLinker) ChangeMedia(m *base.Media) {
	_, _ = s.stmtMap[changeMedia].Query(m.Url, m.Title, m.Uid)
}

func (s *SqlLinker) DeleteMedia(uid int64) error {
	_, err := s.stmtMap[deleteMedia].Query(uid)
	return err
}

func (s *SqlLinker) ChangeArticle(article *base.Article) {
	_, _ = s.stmtMap[updateArticle].Query(article.Summary, article.Id)
}

func (s *SqlLinker) DeleteArticle(key int64) error {
	_, err := s.stmtMap[deleteArticle].Query(key)
	return err
}

func (s *SqlLinker) GetPinsInArea(east, west, north, south float64, timeBegin, timeEnd int64) []int64 {
	r, err := s.stmtMap[getPinsInArea].Query(south, north, west, east, timeBegin, timeEnd)
	if err != nil {
		return nil
	}
	var id int64
	goal := make([]int64, 0, 10)
	for r.Next() {
		err = r.Scan(&id)
		if err != nil {
			return nil
		}
		goal = append(goal, id)
	}
	return goal
}

func (s *SqlLinker) AllRole() ([]string, []int) {
	rs, err := s.stmtMap[allRole].Query()
	if err != nil {
		return nil, nil
	}
	names := make([]string, 0, 10)
	roles := make([]int, 0, 10)
	var name string
	var role int
	for rs.Next() {
		err = rs.Scan(&name, &role)
		if err != nil {
			return nil, nil
		}
		names = append(names, name)
		roles = append(roles, role)
	}
	return names, roles
}
