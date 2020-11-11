package plugin

import (
	config "github.com/ckalpakoglu/rofi-github/config"
	gcache "github.com/ckalpakoglu/rofi-github/cache"
)

type Plugin struct {
	browser string
	conf config.Config
	cache gcache.Cache
}

func New(browser, configFile, cachefile string) Plugin {
	return Plugin{
		browser: browser,
		conf: config.New(configFile),
		cache: gcache.New(cachefile),
	}
}

func (p Plugin) Run(args ...string) error {
	p.cache.Dump()
	return nil
}
