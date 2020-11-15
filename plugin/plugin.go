package plugin

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"

	config "github.com/ckalpakoglu/rofi-github/config"
)

type Plugin struct {
	browser   string
	cacheFile string
	conf      config.Config
}

func New(browser, cachefile string, conf config.Config) Plugin {
	return Plugin{
		browser:   browser,
		cacheFile: cachefile,
		conf:      conf,
	}
}

func (p Plugin) Run(args ...string) error {
	cache := LoadCache(p.cacheFile)

	cmd := exec.Command("rofi", "-dmenu")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func(c Cache) {
		defer stdin.Close()
		for _, val := range c.Content {
			io.WriteString(stdin, fmt.Sprintf("%s\n", val))
		}
	}(cache)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	selected := strings.TrimSpace(string(out))
	if !cache.Exists(selected) {
		r, err := getAllRepoNames(p.conf.User, p.conf.OauthToken)
		if err != nil {
			return err
		}

		for _, v := range r {
			if fuzzy.Match(selected, v.Name) {
				selected = v.HTMLURL
				break
			}
		}
	}

	//if err := cache.Update(selected); err != nil {
	//	return err
	//}

	fmt.Println("====>", selected)
	cmd = exec.Command(p.browser, selected)
	return cmd.Run()
}
