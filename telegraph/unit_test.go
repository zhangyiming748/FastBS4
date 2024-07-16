package ph

import (
	"github.com/zhangyiming748/FastBS4/util"
	"testing"
)

func TestGetFromFile(t *testing.T) {
	GetFromFile()
}
func TestGetFromWeb(t *testing.T) {
	GetFromWeb("https://telegra.ph/%E9%9B%AA%E6%99%B4-%E7%B4%AB%E8%89%B2%E6%81%B6%E9%AD%94%E7%AB%9E%E9%80%9F-05-28")
}

/*
go test -v -run TestGetAllFromWeb
*/
func TestGetAllFromWeb(t *testing.T) {
	urls := util.ReadByLine("urls.txt")
	for _, url := range urls {
		GetFromWeb(url)
	}
}
