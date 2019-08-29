package actions

import (
	"TuriteaWebResources/server/base"
	"TuriteaWebResources/server/buffer"
	"TuriteaWebResources/server/dataLevel"
)

func resourcePreloadOnId(resources []dataLevel.Resource) {
	for _, v := range resources {
		switch v.Type {
		case dataLevel.ArticleContent:
			buffer.MainCache.LoadAsynchronous(dataLevel.ArticleContentKey(v.Id))
		case dataLevel.Image:
		case dataLevel.Video:
			buffer.MainCache.LoadAsynchronous(base.MediaKey(v.Id))
		}
	}
}

func resourcePreloadOnMedia(key dataLevel.ImageKey) {
	buffer.MainCache.LoadAsynchronous(key)
}
