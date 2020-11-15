package plugin

import (
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"os"
	"testing"
)

func TestClientReq(t *testing.T) {
	token := os.Getenv("GH_TOKEN")
	user := os.Getenv("GH_USER")

	r, err := getAllRepoNames(user, token)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range r {
		if fuzzy.Match("twrap", v.Name){
			fmt.Println("===", v)
		}
	}
}
