package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"

	"github.com/ckalpakoglu/rofi-github/plugin"
)

var Version = "v0.0"

func main() {
	home := os.Getenv("HOME")
	var (
		browser  string
		config  string
		cacheFile  string
	)

	app := &cli.App{
		Name: "rofi-github",
		Description: "A Rofi plugin to access github repositories quickly. It does not requires configuration, uses github cli (gh) oauth token",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "browser",
				Aliases:     []string{"b"},
				Value: 		 "google-chrome-stable",
				Usage:       "open in browser",
				Destination: &browser,
			},
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value: 		 home + "/.config/gh/hosts.yml",
				Usage:       "config location",
				Destination: &config,
			},
			&cli.StringFlag{
				Name:        "cache",
				Value: 		 home + "/.config/gh/rofi-github.cache",
				Usage:       "cache file location",
				Destination: &cacheFile,
			},
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "prints version",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("version") {
				return cli.Exit(Version, 0)
			}

			p := plugin.New(browser, config, cacheFile)
			return p.Run()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}