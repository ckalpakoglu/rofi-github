package plugin

import (
	"bufio"
	"log"
	"net/url"
	"os"
	"strings"
)

type Cache struct {
	fname string

	Content []string
}

func LoadCache(fname string) Cache {
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var content []string
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return Cache{
		Content: content,
		fname:   fname,
	}
}

func (c *Cache) Update(selected string) error {
	if _, err := url.Parse(selected); err != nil {
		return err
	}

	f, err := os.OpenFile(c.fname, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	c.Content = append(c.Content, selected+"\n")
	w := bufio.NewWriter(f)
	for i:=0; i <len(c.Content); i++ {
		w.WriteString(c.Content[i]+"\n")
	}

	w.Flush()
	return nil
}

func (c *Cache) Exists(s string) bool {
	for _, v := range c.Content {
		if v == strings.TrimSpace(s) {
			return true
		}
	}
	return false
}
