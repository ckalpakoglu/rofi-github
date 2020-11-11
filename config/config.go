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
		User string `yaml:"user"`
		OatuhToken string `yaml:"oauth_token"`
		GitProtocol string `yaml:"git_protocol"`
	} `yaml:"github.com"`
}

func New(fname string) Config {
	var cfg yamlConf

	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		User: cfg.GithubCom.User,
		OauthToken: cfg.GithubCom.OatuhToken,
		GitProtocol: cfg.GithubCom.GitProtocol,
	}
}
