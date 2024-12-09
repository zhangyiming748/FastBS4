package liwenzhou

import (
	"fmt"
	"testing"
	"time"
)

// go test -v -run TestGetAllFromWeb -timeout 350m
func TestGetAllFromWeb(t *testing.T) {
	items := GetFromArchive()
	for i := 0; i < len(items); i++ {
		GetOne(items[i].Link, items[i].Title)
		for count := 10; count > 0; count-- {
			fmt.Printf("冷却时间还有%d秒\n", count)
			time.Sleep(1 * time.Second)
		}
	}
}

func TestGetOne(t *testing.T) {
	GetOne("https://www.liwenzhou.com/posts/Go/bun", "SQL优先的 Go ORM 框架——Bun 介绍")
}
