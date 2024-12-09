package liwenzhou

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/soup"
	"log"
)

type item struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	//PubDate      string `json:"pubDate"`
	//Guid         string `json:"guid"`
	//Describption string `json:"describ"`
}

func GetFromRss() []item {
	//https://www.liwenzhou.com/index.xml
	uri := "https://www.liwenzhou.com/index.xml"
	resp, err := soup.Get(uri)
	if err != nil {
		log.Panicln(err)
	}
	var is []item
	root := soup.HTMLParse(resp)
	items := root.FindAll("item")
	for i, elem := range items {
		if i == 0 {
			continue
		}
		fmt.Printf("第%d个item:%s\n", i, elem.FullText())
		it := item{
			Title: elem.Find("title").Text(),
			Link:  elem.Find("link").Text(),
		}
		is = append(is, it)
	}
	return is
}
func GetFromArchive() []item {
	//https://www.liwenzhou.com/index.xml
	uri := "https://www.liwenzhou.com/archives/"
	resp, err := soup.Get(uri)
	if err != nil {
		log.Panicln(err)
	}
	var is []item
	root := soup.HTMLParse(resp)
	lis := root.FindAll("li")
	for i, li := range lis {
		a := li.Find("a")
		it := item{
			Title: a.Attrs()["title"],
			Link:  a.Attrs()["href"],
		}
		fmt.Printf("获取到的%d单个结构体%+v\n", i, it)
		is = append(is, it)
	}
	return is
}
