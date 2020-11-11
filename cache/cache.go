package cache

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Cache struct {
	fname string
	content []string
}

func New(fname string) Cache {
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var content []string
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF{
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		content = append(content, line)
	}

	return Cache{
		fname: fname,
		content: content,
	}
}

func (c *Cache) Dump(key ...string) {
	for k, v := range c.content {
		fmt.Print(k, ":", v)
	}
}

