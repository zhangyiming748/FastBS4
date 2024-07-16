package ph

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/zhangyiming748/FastBS4/soup"
	"github.com/zhangyiming748/FastBS4/util"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFromWeb(uri string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("未捕获的错误", err)
		}
	}()
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
	title := root.Find("title").Text()
	log.Printf("标题:%v\n", title)
	os.Mkdir(title, 0777)
	imgs := root.FindAll("img")
	bar := progressbar.New(len(imgs))
	defer bar.Finish()
	for i, img := range imgs {
		bar.Set(i)
		src := img.Attrs()["src"]
		src = strings.Join([]string{"https://telegra.ph", src}, "")
		//fmt.Println(src)
		fname := strings.Join([]string{title, string(os.PathSeparator), strconv.Itoa(i), ".jpg"}, "")
		util.HttpGetProxy("http://192.168.1.5:8889", src, fname)
	}
}
