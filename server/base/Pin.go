package base

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

type Pin struct {
	Uid         int64   `json:"uid"`
	Owner       int64   `json:"owner"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Time        int64   `json:"time"`
	Description string  `json:"description"`
	TagType     string   `json:"tag_type"`
}
var pinPool *sync.Pool
var pinIdChan = make(chan int64, 100)
var pinIdRecycle = make(chan int64, 100)
func init() {
	pinPool = new(sync.Pool)
	pinPool.New = func() interface{} {
		return &Pin{}
	}
	go pinIdProvider()
	if err := loadTags(tagMap); err != nil {
		panic(err)
	}
}

func pinIdProvider() {
	var id int64 = 3
	for {
		select {
		case pinIdChan <- id:
			id ++
		case Id := <-pinIdRecycle:
			pinIdChan<-Id
		}
	}
}
const DefaultTime int64 = 0xffffffff
var tagMap = make(map[uint8]string, 117)
func GenPin(id, owner int64, latitude, longitude float64, t int64, tagType uint8, description string, newOne bool) *Pin {
	if newOne {
		id = <-pinIdChan
	}
	pin := pinPool.Get().(*Pin)
	pin.Uid = id
	pin.Latitude = latitude
	pin.Longitude = longitude
	if t == DefaultTime {
		t = time.Now().Unix()
	}
	pin.Time = t
	pin.TagType = tagMap[tagType]
	pin.Owner = owner
	pin.Description = description
	return pin
}

func RecyclePin(pin *Pin, delete bool) {
	if delete {
		pinIdRecycle <- pin.Uid
	}
	pinPool.Put(pin)
}

func loadTags(m map[uint8]string) error {

	fs, err := ioutil.ReadDir("./cesium/Source/Assets/Textures/maki")
	if err != nil {
		return err
	}
	for i, f := range fs {
		m[uint8(i)] = strings.Split(f.Name(), ".")[0]
	}
	return nil
}

func PinsToJson(pins []*Pin, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(pins)
}

func JsonToPins(r io.Reader, num uint16) ([]*Pin, error) {
	d := json.NewDecoder(r)
	d.UseNumber()
	goal := make([]*Pin, num)
	for i := uint16(0); i < num; i++ {
		goal[i] = pinPool.Get().(*Pin)
	}
	var err error
	err = d.Decode(&goal)
	return goal, err
}