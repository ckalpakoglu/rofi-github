package plugin

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	GithubCom struct {
		User string `yaml:"user"`
		OatuhToken string `yaml:"oauth_token"`
		GitProtocol string `yaml:"git_protocol"`
	} `yaml:"github.com"`
}

func loadConf(fname string) (*Conf, error) {
	c := &Conf{}

	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return c, err
	}

	return c, nil
}

