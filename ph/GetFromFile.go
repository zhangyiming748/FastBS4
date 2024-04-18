package ph

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/soup"
	"os"
	"strings"
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
		if strings.Contains(href, "viewkey=") {
			link := strings.Join([]string{"https://cn.pornhub.com", href}, "")
			fmt.Println(link)
		}
	}
}
