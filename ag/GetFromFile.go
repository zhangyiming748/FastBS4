package ph

import (
	"fmt"
	"os"
	"strings"

	"github.com/zhangyiming748/FastBS4/soup"
)

const (
	host = "https://asmr.121231234.xyz"
)

func GetFromFile(outName string) {
	open, err := os.ReadFile("exam.html")
	if err != nil {
		return
	}
	root := soup.HTMLParse(string(open))
	as := root.FindAll("a")
	file, err := os.OpenFile(outName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
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
