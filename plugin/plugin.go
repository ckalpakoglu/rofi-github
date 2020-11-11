package plugin

import "fmt"

type Plugin struct {
	browser string
	config string
	cachefile string
}

func Run(browser, config, cachefile string) error {
	p := Plugin{
		browser: browser,
		config: config,
		cachefile: cachefile,
	}

	// load config
	conf, err := loadConf(p.config)
	if err != nil {
		return fmt.Errorf("failed to load config file: [%v]", err)
	}
	fmt.Println(conf.GithubCom.User)

	// load cachefile
	cache, err := loadCache(p.cachefile)
	if err != nil {
		return fmt.Errorf("failed to load cache file: [%v]", err)
	}

	fmt.Println("===", p.browser)
	cache.Dump()

	return nil
}
