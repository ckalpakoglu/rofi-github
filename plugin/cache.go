package plugin

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Cache struct {
	content []string
}

func loadCache(fname string) (Cache, error) {
	var c Cache

	//f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f, err := os.Open(fname)
	if err != nil {
		return c, err
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return c, err
		}
		c.content = append(c.content, line)
	}

	return c, nil

}

func (c *Cache) Dump(key ...string) {
	for k, v := range c.content {
		fmt.Print(k, ":", v)
	}
}

