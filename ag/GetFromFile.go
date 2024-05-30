package ph

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/soup"
	"os"
	"strings"
)

const (
	host = "https://www.asmrgay.com"
)

func GetFromFile() {
	open, err := os.ReadFile("exam.html")
	if err != nil {
		return
	}
	root := soup.HTMLParse(string(open))
	as := root.FindAll("a")
	for _, a := range as {
		href := a.Attrs()["href"]
		if strings.HasPrefix(href, "/asmr/中文音声") {
			href = strings.Join([]string{host, href}, "")
			fmt.Println(href)
		}

	}
}
