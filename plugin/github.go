package plugin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type repo struct {
	Name     string `json:"name"`
	HTMLURL  string `json:"html_url"`
	CloneURL string `json:"clone_url"`
}

func getAllRepoNames(user, token string) ([]repo, error) {
	var all []repo
	var client http.Client
	var url string
	{
		endpoint := "https://api.github.com/orgs"
		url = fmt.Sprintf("%s/%s/repos?per_page=1000", endpoint, user)
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := client.Do(req)
	if err != nil {
		return all, err

	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return all, err
	}

	if err = json.Unmarshal(data, &all); err !=  nil{
		return all, err
	}

	return all, nil
}
