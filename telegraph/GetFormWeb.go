package ph

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/soup"
	"os"
	"strings"
)

func GetFromWeb(uri string) {
	resp, err := soup.GetWithProxy(uri, "http://192.168.1.5:8889")
	root := soup.HTMLParse(resp)
	if err != nil {
		fmt.Println("get 出现错误", err)
	}
	file, err := os.OpenFile("exam.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(resp)
	imgs := root.FindAll("img")
	for _, img := range imgs {
		src := img.Attrs()["src"]
		src = strings.Join([]string{"https://telegra.ph", src}, "")
		fmt.Println(src)
	}
}
