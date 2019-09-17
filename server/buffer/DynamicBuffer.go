package buffer

import (
	"log"
	"sync"
	"time"

	"github.com/ChenXingyuChina/asynchronousIO"

	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/dataLevel"
)

type cacheBlock struct {
	lock *sync.Mutex
	cache map[uint64]*item
}

const (
	update = 1 << iota
	deleted
	notExist
	normal
)

var MainCache = NewCache()


type item struct {
	data asynchronousIO.Bean
	updateState uint8
	lastModifyTime int64
}

// no use but may be useful
//type aKey struct {
//	t int64
//	uid int64
//}

type Cache struct {
	caches [256]cacheBlock
	itemPool *sync.Pool
}

func (c *Cache) Clear() {
	ticker := time.NewTicker(4*time.Minute >> 8)
	clock := ticker.C
	number := uint8(0)
	for {
		t := <-clock
		block := c.caches[number]
		block.lock.Lock()
		if uint64(len(block.cache)) & 0xfffffffffffffff0 != 0 {
			block.clear(number, t.Unix())
		}
		block.lock.Unlock()
		number++
	}
}

func (c *cacheBlock) clear(blockNumber uint8, t int64) {
	saveList := make([]asynchronousIO.Bean, 0, 20)
	recycleList := make([]asynchronousIO.Bean, 0, 30)
	for k, o := range c.cache {
		if t - o.lastModifyTime < int64(5 * time.Minute) {
			continue
		}
		switch o.updateState {
		case update:
			delete(c.cache, k)
			saveList = append(saveList, o.data)
		case notExist:
		case deleted:
			delete(c.cache, k)
		case normal:
			delete(c.cache, k)
			recycleList = append(recycleList, o.data)
		}

	}
	go func(saveList []asynchronousIO.Bean, recycleList []asynchronousIO.Bean) {
		callbackList := make([]func() error, len(saveList))
		for i, bean := range saveList {
			f := dataLevel.SaveAndNotify(bean)
			callbackList[i] = f
		}
		for _, v := range recycleList {
			dataLevel.RecycleData(v)
		}
		for i, f := range callbackList {
			err := f()
			if err != nil {
				//todo handle errors and try to recover
				log.Printf("%s: %v", time.Now().Format(time.Stamp), err)
			}
			dataLevel.RecycleData(saveList[i])
		}
	}(saveList, recycleList)
}

func (c *Cache) Load(key asynchronousIO.Key) (goal asynchronousIO.Bean, ok bool) {
	t := key.TypeId()
	uid, _ := key.UniqueId()
	b := uint8(uid) + uint8(t)
	u := (uint64(uid) >> 8 << 8) | uint64(t)
	cache := c.caches[b]
	cache.lock.Lock()
	if v, exist := cache.cache[u]; exist {
		v.lastModifyTime = time.Now().Unix()
		if v.updateState & (deleted | notExist) != 0 {
			cache.lock.Unlock()
			return nil, false
		}
		cache.lock.Unlock()
		return v.data, true
	}
	cache.lock.Unlock()
	f := dataLevel.Load(key)
	goal, err := f()
	if err == nil {
		i := c.itemPool.Get().(*item)
		i.data = goal
		i.updateState = normal
		cache.lock.Lock()
		if in, ok := cache.cache[u]; ok {
			goal = in.data
			cache.lock.Unlock()
			dataLevel.RecycleData(i.data)
			c.itemPool.Put(i)
			return goal, true
		} else{
			i.lastModifyTime = time.Now().Unix()
			cache.cache[u] = i
			cache.lock.Unlock()
			return goal, true
		}
	} else {
		i := c.itemPool.Get().(*item)
		i.lastModifyTime = time.Now().Unix()
		i.updateState = notExist
		cache.lock.Lock()
		cache.cache[u] = i
		cache.lock.Unlock()
		return nil, false
	}
}

func (c *Cache) LoadAsynchronous(key asynchronousIO.Key) {
	t := key.TypeId()
	uid, _ := key.UniqueId()
	b := uint8(uid) + uint8(t)
	u := (uint64(uid) >> 8 << 8) | uint64(t)
	cache := c.caches[b]
	cache.lock.Lock()
	if v, exist := cache.cache[u]; exist {
		v.lastModifyTime = time.Now().Unix()
		if v.updateState & (deleted | notExist) != 0 {
			cache.lock.Unlock()
			return
		}
		cache.lock.Unlock()
		return
	}
	cache.lock.Unlock()
	go c.loadAsynchronousHelp(key, cache, u)
}

func (c *Cache) loadAsynchronousHelp(key asynchronousIO.Key, cache cacheBlock, idInBlock uint64) {
	f := dataLevel.Load(key)
	goal, err := f()
	if err == nil {
		i := c.itemPool.Get().(*item)
		i.data = goal
		i.updateState = normal
		cache.lock.Lock()
		if _, ok := cache.cache[idInBlock]; ok {
			cache.lock.Unlock()
			dataLevel.RecycleData(i.data)
			c.itemPool.Put(i)
		} else{
			i.lastModifyTime = time.Now().Unix()
			cache.cache[idInBlock] = i
			cache.lock.Unlock()
		}
	} else {
		i := c.itemPool.Get().(*item)
		i.lastModifyTime = time.Now().Unix()
		i.updateState = notExist
		cache.lock.Lock()
		cache.cache[idInBlock] = i
		cache.lock.Unlock()
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
		v.lastModifyTime = time.Now().Unix()
		if v.updateState & (deleted | notExist) != 0 {
			cache.lock.Unlock()
			return true
		}
		data = v.data
		v.data = nil
		v.updateState = deleted
		cache.lock.Unlock()
		f := dataLevel.Delete(key)
		err := f()
		if err != nil {
			switch key.(type) {
			case base.ArticleKey:
				return false
			case base.PinKey:
				return false
			case base.MediaKey:
				return false
			}
		}
		dataLevel.RecycleData(data)
		return true
	}
	v := c.itemPool.Get().(*item)
	v.updateState = deleted
	v.data = nil
	v.lastModifyTime = time.Now().Unix()
	cache.lock.Unlock()
	f := dataLevel.Delete(key)
	//fmt.Print(f)
	err := f()
	if err != nil {
		//fmt.Println(err)
		return false
	}
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
		v.lastModifyTime = time.Now().Unix()
		v.updateState = update
		oldData := v.data
		v.data = bean
		cache.lock.Unlock()
		dataLevel.RecycleData(oldData)
		return
	}
	item := c.itemPool.Get().(*item)
	item.data = bean
	item.updateState = update
	item.lastModifyTime = time.Now().Unix()
	cache.cache[u] = item
	cache.lock.Unlock()
}

func (c *Cache) CreatePin(pin *base.Pin) bool {
	b := uint8(pin.Uid) + uint8(dataLevel.Pin)
	u := (uint64(pin.Uid) >> 8 << 8) | uint64(dataLevel.Pin)
	cache := c.caches[b]
	ok := dataLevel.SQLWorker.CreatePin(pin.Uid, pin.Owner, pin.Latitude, pin.Longitude, pin.Time, base.TagNameToNumber[pin.TagType], pin.Description, pin.Name, pin.Color)
	if ok {
		item := c.itemPool.Get().(*item)
		item.updateState = normal
		item.data = pin
		item.lastModifyTime =time.Now().Unix()
		cache.lock.Lock()
		cache.cache[u] = item
		cache.lock.Unlock()
	} else {
		base.RecyclePin(pin, true)
	}
	return ok
}

func (c *Cache) CreateArticle(article *base.Article) bool {
	b := uint8(article.Id) + uint8(dataLevel.Article)
	u := (uint64(article.Id) >> 8 << 8) | uint64(dataLevel.Article)
	cache := c.caches[b]
	ok := dataLevel.SQLWorker.CreateArticle(article.Summary, article.Id, article.WriteBy, article.HomeContent)
	if ok {
		item := c.itemPool.Get().(*item)
		item.updateState = normal
		item.data = article
		item.lastModifyTime =time.Now().Unix()
		cache.lock.Lock()
		cache.cache[u] = item
		cache.lock.Unlock()
	} else {
		base.RecycleArticle(article, true)
	}
	return ok
}

func (c *Cache) CreateMedia(media *base.Media) bool {
	b := uint8(media.Uid) + uint8(dataLevel.Media)
	u := (uint64(media.Uid) >> 8 << 8) | uint64(dataLevel.Media)
	cache := c.caches[b]
	ok := dataLevel.SQLWorker.AddMedia(media.Uid, media.Title, media.Url, media.Type)
	if ok {
		item := c.itemPool.Get().(*item)
		item.updateState = normal
		item.data = media
		item.lastModifyTime =time.Now().Unix()
		cache.lock.Lock()
		cache.cache[u] = item
		cache.lock.Unlock()
	} else {
		base.RecycleMedia(media, true)
	}
	return ok
}

func (c *Cache) CreateArticleContent(resources []dataLevel.Resource, content string) int64 {
	ac := dataLevel.CreateContentByData(resources, content)
	b := uint8(ac.Id) + uint8(dataLevel.ArticleContentResources)
	u := (uint64(ac.Id) >> 8 << 8) | uint64(dataLevel.ArticleContentResources)
	cache := c.caches[b]
	item := c.itemPool.Get().(*item)
	item.updateState = update
	item.data = ac
	item.lastModifyTime =time.Now().Unix()
	cache.lock.Lock()
	cache.cache[u] = item
	cache.lock.Unlock()
	return ac.Id
}

func (c *Cache) CreateImage(data []byte, id int64){
	image := dataLevel.CreateImageByData(data)
	image.Id = id
	b := uint8(image.Id) + uint8(dataLevel.ImagesResources)
	u := (uint64(image.Id) >> 8 << 8) | uint64(dataLevel.ImagesResources)
	cache := c.caches[b]
	item := c.itemPool.Get().(*item)
	item.updateState = update
	item.data = image
	item.lastModifyTime =time.Now().Unix()
	cache.lock.Lock()
	cache.cache[u] = item
	cache.lock.Unlock()
}

// maybe just for test
func (c *Cache) flushBlock(number uint8) {
	block := c.caches[number]
	block.lock.Lock()
	block.clear(number, 0x7fffffffffffffff)
	block.lock.Unlock()
}

func NewCache() *Cache {
	goal := &Cache{}
	go goal.Clear()
	for i := 0; i < 256; i++ {
		goal.caches[i].lock = new(sync.Mutex)
		goal.caches[i].cache = make(map[uint64]*item)
	}
	goal.itemPool = new(sync.Pool)
	goal.itemPool.New = func() interface{} {
		return &item{}
	}
	return goal
}

const (
	Exist = 1 << iota
	Deleted
	NotExist
	NotInBuffer
)
func (c *Cache) LoadIfExist(key asynchronousIO.Key) (asynchronousIO.Bean, uint8) {
	t := key.TypeId()
	uid, _ := key.UniqueId()
	b := uint8(uid) + uint8(t)
	u := (uint64(uid) >> 8 << 8) | uint64(t)
	cache := c.caches[b]
	cache.lock.Lock()
	if v, exist := cache.cache[u]; exist {
		v.lastModifyTime = time.Now().Unix()
		if v.updateState & (deleted | notExist) != 0 {
			cache.lock.Unlock()
			return nil, v.updateState
		} else if v.updateState == normal {
			cache.lock.Unlock()
			return v.data, Exist
		}
		cache.lock.Unlock()
		return v.data, v.updateState
	}
	cache.lock.Unlock()
	return nil, NotInBuffer
}
