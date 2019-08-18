package buffer

import (
	"TuriteaWebResources/asynchronousIO"
	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/dataLevel"
	"log"
	"sort"
	"sync"
	"time"
)

type cacheBlock struct {
	lock *sync.RWMutex
	cache map[uint64]*item
}

const (
	create = 1 << iota
	createButUpdate
	update
	deleted
	notExist
	normal
)



type item struct {
	data asynchronousIO.Bean
	updateState uint8
}

type key struct {
	t uint8
	uid int64
}

type Cache struct {
	lastTimeTable map[key]int64
	caches [256]cacheBlock
	lastTimeChan chan key
	itemPool *sync.Pool
}

func (c *Cache) statistics() {
	for {
		k := <- c.lastTimeChan
		c.lastTimeTable[k] = time.Now().Unix()
		if len(c.lastTimeTable) > 5000{
			c.clear()
		}
	}
}

type helperForSortMap struct {
	keys []key
	values []int64
}

func (h *helperForSortMap) Len() int {
	return len(h.values)
}

// 升序
func (h *helperForSortMap) Less(i, j int) bool {
	return h.values[i] < h.values[j]
}

func (h *helperForSortMap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
	h.keys[i], h.keys[j] = h.keys[j], h.keys[i]
}

func fromMapToHelp(m map[key]int64) *helperForSortMap {
	goal := &helperForSortMap{
		make([]key, len(m)),
		make([]int64, len(m)),
	}
	var i int
	for k, v := range m {
		goal.keys[i] = k
		goal.values[i] = v
		i++
	}
	return goal
}

func (c *Cache) clear() {
	h := fromMapToHelp(c.lastTimeTable)
	sort.Sort(h)
	c.lastTimeTable = make(map[key]int64)
	remove := 0
	callback := make([]func() error, 0, 500)
	recycleList := make([]asynchronousIO.Bean, 0, 700)
	for i, v := range h.keys {
		if remove < 1000 {
			b := uint8(v.uid) + v.t
			u := (uint64(v.uid) >> 8 << 8) | uint64(v.t)
			cache := c.caches[b]
			cache.lock.Lock()
			o := c.caches[b].cache[u]
			switch o.updateState {
			case update:
				delete(cache.cache, u)
				cache.lock.Unlock()
				remove++
				f := dataLevel.SaveAndNotify(o.data)
				callback = append(callback, f)
				recycleList = append(recycleList, o.data)
			case notExist:
			case deleted:
				delete(cache.cache, u)
				cache.lock.Unlock()
				remove ++
			case create:
			case createButUpdate:
				c.lastTimeTable[v] = h.values[i]
				cache.lock.Unlock()
			case normal:
				delete(cache.cache, u)
				cache.lock.Unlock()
				remove++
				recycleList = append(recycleList, o.data)
			}
		} else {
			c.lastTimeTable[v] = h.values[i]
		}
	}
	go func(callback []func() error, recycleList []asynchronousIO.Bean) {
		for _, f := range callback {
			err := f()
			if err != nil {
				//todo handle errors and try to recover
				log.Printf("%s: %v", time.Now().Format(time.Stamp), err)
			}
		}
		for _, v := range recycleList {
			dataLevel.RecycleData(v)
		}
	}(callback, recycleList)
}

func (c *Cache) Load(key asynchronousIO.Key) (goal asynchronousIO.Bean, ok bool) {
	t := key.TypeId()
	uid, _ := key.UniqueId()
	b := uint8(uid) + uint8(t)
	u := (uint64(uid) >> 8 << 8) | uint64(t)
	cache := c.caches[b]
	cache.lock.RLock()
	if v, exist := cache.cache[u]; exist {
		if v.updateState & (deleted | notExist) != 0 {
			cache.lock.RUnlock()
			return nil, false
		}
		cache.lock.RUnlock()
		return v.data, true
	}
	cache.lock.RUnlock()
	f := dataLevel.Load(key)
	goal, err := f()
	if err != nil {
		i := c.itemPool.Get().(*item)
		i.data = goal
		i.updateState = normal
		cache.lock.Lock()
		if _, ok := cache.cache[u]; !ok {
			cache.lock.Unlock()
			dataLevel.RecycleData(i.data)
			c.itemPool.Put(i)
			return goal, true
		} else{
			cache.cache[u] = i
			cache.lock.Unlock()
			return goal, true
		}
	} else {
		return nil, false
	}
}

func (c *Cache) LoadAsynchronous(key asynchronousIO.Key) {
	t := key.TypeId()
	uid, _ := key.UniqueId()
	b := uint8(uid) + uint8(t)
	u := (uint64(uid) >> 8 << 8) | uint64(t)
	cache := c.caches[b]
	cache.lock.RLock()
	if v, exist := cache.cache[u]; exist {
		if v.updateState & (deleted | notExist) != 0 {
			cache.lock.RUnlock()
			return
		}
		cache.lock.RUnlock()
		return
	}
	cache.lock.RUnlock()
	go c.loadAsynchronousHelp(key, cache, u)
}

func (c *Cache) loadAsynchronousHelp(key asynchronousIO.Key, cache cacheBlock, idInBlock uint64) {
	f := dataLevel.Load(key)
	goal, err := f()
	if err != nil {
		i := c.itemPool.Get().(*item)
		i.data = goal
		i.updateState = normal
		cache.lock.Lock()
		if _, ok := cache.cache[idInBlock]; !ok {
			cache.lock.Unlock()
			dataLevel.RecycleData(i.data)
			c.itemPool.Put(i)
		} else{
			cache.cache[idInBlock] = i
			cache.lock.Unlock()
		}
	}
}

func (c *Cache) Delete(key asynchronousIO.Key) bool {
	t := key.TypeId()
	uid, _ := key.UniqueId()
	b := uint8(uid) + uint8(t)
	u := (uint64(uid) >> 8 << 8) | uint64(t)
	cache := c.caches[b]
	var data asynchronousIO.Bean
	cache.lock.Lock()
	if v, exist := cache.cache[u]; exist {
		if v.updateState & (deleted | notExist) != 0 {
			cache.lock.Unlock()
			return true
		}
		data = v.data
		v.updateState = deleted
		cache.lock.Unlock()
		f := dataLevel.Delete(key)
		err := f()
		if err != nil {
			return false
		}
		cache.lock.Lock()
		if v, exist := cache.cache[u]; exist && v.updateState == deleted{
			dataLevel.RecycleData(data)
			v.data = nil
		}
		cache.lock.Unlock()
		return true
	}
	v := cache.cache[u]
	v.updateState = deleted
	v.data = nil
	cache.lock.Unlock()
	f := dataLevel.SaveAndNotify(v.data)
	err := f()
	if err != nil {
		return false
	}
	dataLevel.RecycleData(v.data)
	return true
}

func (c *Cache) Update(bean asynchronousIO.Bean) {
	key := bean.GetKey()
	t := key.TypeId()
	uid, _ := key.UniqueId()
	b := uint8(uid) + uint8(t)
	u := (uint64(uid) >> 8 << 8) | uint64(t)
	cache := c.caches[b]
	cache.lock.Lock()
	if v, exist := cache.cache[u]; exist {
		if v.updateState == create {
			v.updateState = createButUpdate
			cache.lock.Unlock()
			return
		}
		v.updateState = update
		oldData := v.data
		v.data = bean
		cache.lock.Unlock()
		dataLevel.RecycleData(oldData)
	}
	item := c.itemPool.Get().(*item)
	item.data = bean
	item.updateState = update
	cache.cache[u] = item
	cache.lock.Unlock()
}

func (c *Cache) CreatePin(owner int64, latitude, longitude float64, t int64, tagType uint8, description, name string) bool {
	pin := base.GenPin(0, owner, latitude, longitude, t, tagType, description, name, true)
	b := uint8(pin.Uid) + uint8(dataLevel.Pin)
	u := (uint64(pin.Uid) >> 8 << 8) | uint64(dataLevel.Pin)
	cache := c.caches[b]
	item := c.itemPool.Get().(*item)
	item.updateState = create
	item.data = pin
	cache.lock.Lock()
	cache.cache[u] = item
	cache.lock.Unlock()
	ok := dataLevel.SQLWorker.CreatePin(pin.Uid, owner, latitude, longitude, t, tagType, description, name)
	if !ok {
		cache.lock.Lock()
		delete(cache.cache, u)
		cache.lock.Unlock()
		base.RecyclePin(item.data.(*base.Pin), true)
		c.itemPool.Put(item)
	}
	return ok
}

func (c *Cache) CreateArticle(writeBy int64, summary string) bool {
	article := base.GenArticle(0, writeBy, summary, true)
	b := uint8(article.Id) + uint8(dataLevel.Article)
	u := (uint64(article.Id) >> 8 << 8) | uint64(dataLevel.Article)
	cache := c.caches[b]
	item := c.itemPool.Get().(*item)
	item.updateState = create
	item.data = article
	cache.lock.Lock()
	cache.cache[u] = item
	cache.lock.Unlock()
	ok := dataLevel.SQLWorker.CreateArticle(summary, article.Id, writeBy)
	if !ok {
		cache.lock.Lock()
		delete(cache.cache, u)
		cache.lock.Unlock()
		base.RecycleArticle(item.data.(*base.Article), true)
		c.itemPool.Put(item)
	}
	return ok
}

func (c *Cache) CreateMedia(title, url string, t uint8) bool {
	media := base.GenMedia(0, t, title, url, true)
	b := uint8(media.Uid) + uint8(dataLevel.Media)
	u := (uint64(media.Uid) >> 8 << 8) | uint64(dataLevel.Media)
	cache := c.caches[b]
	item := c.itemPool.Get().(*item)
	item.updateState = create
	item.data = media
	cache.lock.Lock()
	cache.cache[u] = item
	cache.lock.Unlock()
	ok := dataLevel.SQLWorker.AddMedia(media.Uid, title, url, t)
	if !ok {
		cache.lock.Lock()
		delete(cache.cache, u)
		cache.lock.Unlock()
		base.RecycleMedia(item.data.(*base.Media), true)
		c.itemPool.Put(item)
	}
	return ok
}

func (c *Cache) CreateArticleContent(resources []dataLevel.Resource, content string) {
	ac := dataLevel.CreateContentByData(resources, content)
	b := uint8(ac.Id) + uint8(dataLevel.Article)
	u := (uint64(ac.Id) >> 8 << 8) | uint64(dataLevel.Article)
	cache := c.caches[b]
	item := c.itemPool.Get().(*item)
	item.updateState = update
	item.data = ac
	cache.lock.Lock()
	cache.cache[u] = item
	cache.lock.Unlock()
}

func (c *Cache) CreateImage(data []byte) {
	image := dataLevel.CreateImageByData(data)
	b := uint8(image.Id) + uint8(dataLevel.Article)
	u := (uint64(image.Id) >> 8 << 8) | uint64(dataLevel.Article)
	cache := c.caches[b]
	item := c.itemPool.Get().(*item)
	item.updateState = update
	item.data = image
	cache.lock.Lock()
	cache.cache[u] = item
	cache.lock.Unlock()
}

func NewCache() *Cache {
	goal := &Cache{lastTimeTable:make(map[key]int64), lastTimeChan:make(chan key, 10000)}
	go goal.statistics()
	for i := 0; i < 256; i++ {
		goal.caches[i].lock = new(sync.RWMutex)
		goal.caches[i].cache = make(map[uint64]*item)
	}
	goal.itemPool = new(sync.Pool)
	goal.itemPool.New = func() interface{} {
		return &item{}
	}
	return goal
}
