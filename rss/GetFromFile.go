package rss

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/soup"
	"github.com/zhangyiming748/FastBS4/util"
	"os"
	"strings"
)

func GetFromFile() {
	open, err := os.ReadFile("rss.xml")
	lines := []string{"#!/usr/bin/env bash"}
	if err != nil {
		return
	}
	root := soup.HTMLParse(string(open))
	fmt.Println(1, root)
	items := root.FindAll("item")
	for i, item := range items {
		//fmt.Println(i, item)
		title := item.Find("itunes:title").Text()
		link := item.Find("enclosure").Attrs()["url"]
		//fmt.Println(title, link)

		name := strings.ReplaceAll(title, "\u300A", "")
		name = strings.ReplaceAll(title, "\u300B", "")
		name = strings.ReplaceAll(title, "\uE000", "")
		name = strings.Join([]string{name, "mp3"}, ".")
		//wget -e use_proxy=yes -e http_proxy=192.168.1.20:8889 -e  https_proxy=192.168.1.20:8889 --no-check-certificate -O
		line := strings.Join([]string{"wget -e use_proxy=yes -e http_proxy=192.168.1.20:8889 -e  https_proxy=192.168.1.20:8889 --no-check-certificate -O", name, link}, " ")
		fmt.Println(i, line)
		lines = append(lines, line)
	}
	util.WriteByLine("down.sh", lines)
}

func GetFromFile2() {
	open, err := os.ReadFile("rss.xml")
	lines := []string{"#!/usr/bin/env bash"}
	if err != nil {
		return
	}
	root := soup.HTMLParse(string(open))
	fmt.Println(1, root)
	items := root.FindAll("item")
	for i, item := range items {
		fmt.Println(i, item)
		title := item.Find("title").Text()
		title = strings.ReplaceAll(title, " ", "")
		title = strings.ReplaceAll(title, "\n", "")
		title = strings.ReplaceAll(title, "<![CDATA[", "")
		title = strings.ReplaceAll(title, "]]>", "")
		title = strings.ReplaceAll(title, "《", "")
		title = strings.ReplaceAll(title, "》", "")
		title = strings.ReplaceAll(title, "（", "")
		title = strings.ReplaceAll(title, "）", "")
		title = strings.ReplaceAll(title, "、", "")
		title = strings.ReplaceAll(title, "从乞丐到元首", "")

		if len(title) == 1 {
			title = "00" + title
		}
		if len(title) == 2 {
			title = "0" + title
		}
		//re := regexp.MustCompile(`\d+\.`)
		//title = re.ReplaceAllString(title, "")

		fmt.Printf("title:%s\n", title)
		name := strings.Join([]string{title, "m4a"}, ".")
		fmt.Println(i, name)
		link := item.Find("enclosure").Attrs()["url"]
		fmt.Println(i, link)
		link = strings.Split(link, "&jt=")[1]
		line := strings.Join([]string{"wget -e use_proxy=yes -e http_proxy=192.168.1.20:8889 -e  https_proxy=192.168.1.20:8889 --no-check-certificate -O", name, link}, " ")
		lines = append(lines, line)
	}
	util.WriteByLine("down.sh", lines)
}
