package ph

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/soup"
	"os"
	"strings"
)

const (
	host = "https://asmr.121231234.xyz"
)

func GetFromFile() {
	open, err := os.ReadFile("exam.html")
	if err != nil {
		return
	}
	root := soup.HTMLParse(string(open))
	as := root.FindAll("a")
	file, err := os.OpenFile("第二季601-700.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	for _, a := range as {
		href := a.Attrs()["href"]
		if strings.HasPrefix(href, "/asmr/中文音声") {
			href = strings.Join([]string{host, href}, "")
			fmt.Println(href)
			file.WriteString(fmt.Sprintf("%s\n", href))
		}

	}
}
