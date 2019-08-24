package actions

import "TuriteaWebResources/server/dataLevel"

func ConfigAction(config Config) {
	if config.PreLoadOnIn {
		dataLevel.OnLoadResourceId = resourcePreloadOnId
	}
	if config.PreLoadOnMedia {
		dataLevel.OnLoadMedia = resourcePreloadOnMedia
	}
	dataLevel.Init()
}

type Config struct {
	PreLoadOnIn bool
	PreLoadOnMedia bool
}
