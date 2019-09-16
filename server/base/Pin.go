package base

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/ChenXingyuChina/asynchronousIO"
)

type Pin struct {
	Uid         int64   `json:"uid"`
	Owner       int64   `json:"owner"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Time        int64   `json:"time"`
	Description string  `json:"description"`
	TagType     string   `json:"tag_type"`
	Name string `json:"name"`
	Color string `json:"color"`
}

func (p *Pin) GetKey() asynchronousIO.Key {
	return PinKey(p.Uid)
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
	if err := loadTags(tagMap, TagNameToNumber); err != nil {
		panic(err)
	}
}

func pinIdProvider() {
	var id int64 = 27
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
var tagMap = [117]string{}
var TagNameToNumber = make(map[string]uint8, 117)
func GenPin(id, owner int64, latitude, longitude float64, t int64, tagType uint8, description, name string, color string) *Pin {
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
	pin.Name = name
	pin.Color = color
	return pin
}

func RecyclePin(pin *Pin, delete bool) {
	if delete {
		pinIdRecycle <- pin.Uid
	}
	pinPool.Put(pin)
}

func loadTags(m [117]string, rm map[string]uint8) error {
	//for test open this one and close other one
	fs, err := ioutil.ReadDir("../../cesium/Source/Assets/Textures/maki")
	//for overallTesting open this
	//fs, err := ioutil.ReadDir("../cesium/Source/Assets/Textures/maki")
	//fs, err := ioutil.ReadDir("./cesium/Source/Assets/Textures/maki")
	if err != nil {
		return err
	}
	for i, f := range fs {
		s := strings.Split(f.Name(), ".")[0]
		m[uint8(i)] = s
		rm[s] = uint8(i)
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

func GenPinId() int64 {
	return <-pinIdChan
}

type PinKey int64

func (p PinKey) UniqueId() (int64, bool) {
	return int64(p), true
}

func (PinKey) ToString() (string, bool) {
	panic("implement me")
}

func (PinKey) TypeId() int64 {
	return 2
}