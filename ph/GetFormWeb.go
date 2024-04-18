package ph

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/soup"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetAllFromWeb(base string) {
	for i := 1; i <= 9; i++ {
		if i == 1 {
			uri := strings.Join([]string{base, "/videos"}, "")
			GetFromWeb(uri)
			fmt.Println(uri)
			continue
		}
		uri := strings.Join([]string{base, "/videos?page=", strconv.Itoa(i)}, "")
		fmt.Println(uri)
		GetFromWeb(uri)
		time.Sleep(3 * time.Second)
	}
}
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

	links, _ := os.OpenFile("links.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer links.Close()
	as := root.FindAll("a")
	for _, a := range as {
		href := a.Attrs()["href"]
		if strings.Contains(href, "viewkey=") {
			link := strings.Join([]string{"https://cn.github.com", href}, "")
			fmt.Println(link)
			links.WriteString(fmt.Sprintf("%s\n", link))
		}
	}
}
