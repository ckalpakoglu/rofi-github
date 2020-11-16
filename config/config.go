package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	User        string
	OauthToken  string
	GitProtocol string
}

type yamlConf struct {
	GithubCom struct {
		User        string `yaml:"user"`
		OatuhToken  string `yaml:"oauth_token"`
		GitProtocol string `yaml:"git_protocol"`
	} `yaml:"github.com"`
}

func Init(fname, org string) Config {
	var cfg yamlConf

	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var user string
	if len(org) > 0 {
		user = org
	} else {
		user = cfg.GithubCom.User
	}

	return Config{
		User:        user,
		OauthToken:  cfg.GithubCom.OatuhToken,
		GitProtocol: cfg.GithubCom.GitProtocol,
	}
}
