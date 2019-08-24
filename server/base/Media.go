package base

import (
	"github.com/ChenXingyuChina/asynchronousIO"
	"encoding/json"
	"io"
	"sync"
)

type Media struct {
	Type uint8 `json:"type"`
	Title string `json:"title"`
	Url string `json:"url"`
	Uid int64 `json:"uid"`
}

func (m *Media) GetKey() asynchronousIO.Key {
	return MediaKey(m.Uid)
}

var mediaPool *sync.Pool
var mediaIdChan = make(chan int64, 100)
var mediaIdRecycle = make(chan int64, 100)
func init() {
	mediaPool = new(sync.Pool)
	mediaPool.New = func() interface{} {
		return &Media{}
	}
	go mediaIdProvider()
}

func mediaIdProvider() {
	var id int64 = 2
	for {
		select {
		case mediaIdChan <- id:
			id ++
		case Id := <-mediaIdRecycle:
			mediaIdChan <- Id
		}
	}
}

func GenMedia(uid int64, t uint8, title string, url string) *Media {
	media := mediaPool.Get().(*Media)
	media.Uid = uid
	media.Type = t
	media.Title = title
	media.Url = url
	return media
}

func RecycleMedia(media *Media, delete bool) {
	if delete {
		mediaIdRecycle <- media.Uid
	}
	mediaPool.Put(media)
}

func GenMediaId() int64 {
	return <-mediaIdChan
}

type MediaKey int64

func (m MediaKey) UniqueId() (int64, bool) {
	return int64(m), true
}

func (MediaKey) ToString() (string, bool) {
	panic("implement me")
}

func (MediaKey) TypeId() int64 {
	return 3
}


func MediaToJson(medias []*Media, w io.Writer) error {
	return json.NewEncoder(w).Encode(medias)
}

func JsonToMedia(r io.Reader, num uint16) ([]*Media, error) {
	d := json.NewDecoder(r)
	d.UseNumber()
	goal := make([]*Media, num)
	for i := uint16(0); i < num; i++ {
		goal[i] = mediaPool.Get().(*Media)
	}
	var err error
	err = d.Decode(&goal)
	return goal, err
}
