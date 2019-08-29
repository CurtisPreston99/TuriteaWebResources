package actions

import (
	"log"
	"net/http"
	"strconv"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func GetPins(w http.ResponseWriter, r *http.Request) {
	log.Println("callGetPins")
	q := r.URL.Query()
	north, err := strconv.ParseFloat(q.Get("north"), 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	south, err := strconv.ParseFloat(q.Get("south"), 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	west, err := strconv.ParseFloat(q.Get("west"), 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	east, err := strconv.ParseFloat(q.Get("east"), 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	timeBegin, err := strconv.ParseInt(q.Get("timeBegin"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	timeEnd, err := strconv.ParseInt(q.Get("timeEnd"), 16, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	pins := dataLevel.SQLWorker.GetPinsInArea(east, west, north, south, timeBegin, timeEnd)
	goal := make([]*base.Pin, len(pins))
	for i, v := range pins {
		bean, ok := buffer.MainCache.Load(base.PinKey(v))
		if !ok {
			goal[i]= nil
		} else {
			goal[i] = bean.(*base.Pin)
		}
	}
	err = base.PinsToJson(goal, w)
	if err != nil {
		w.WriteHeader(400)
		return
	}
}

func GetPin(w http.ResponseWriter, r *http.Request) {
	log.Println("call get pin")
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 16, 64)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	pin, ok := buffer.MainCache.Load(base.PinKey(id))
	if !ok {
		w.WriteHeader(404)
	}
	err = base.PinsToJson([]*base.Pin{pin.(*base.Pin)}, w)
	if err != nil {
		w.WriteHeader(404)
		return
	}
}
