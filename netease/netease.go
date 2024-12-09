package netease

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/soup"
	"log"
	"os"
	"strings"
)

func GetTop200Music(uri string) (ids []string) {
	html, err := soup.Get(uri)
	if err != nil {
		panic(err)
	}
	fmt.Println(html)
	root := soup.HTMLParse(html)
	as := root.FindAll("a")
	shell, _ := os.OpenFile("down.sh", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	for i, a := range as {
		href := a.Attrs()["href"]
		name := a.Text()
		if strings.Contains(href, "/song?id=") {
			id := strings.Split(href, "=")[1]
			log.Printf("第%v首歌%v的id为%v", i+1, name, id)
			prefix := "http://music.163.com/song/media/outer/url?id="
			song := strings.Join([]string{prefix, id}, "")
			name = strings.Join([]string{name, "mp3"}, ".")
			name = strings.Join([]string{"\"", name, "\""}, "")
			wget := strings.Join([]string{"wget", song, "-O", name}, " ")
			shell.WriteString(fmt.Sprintf("%v\n", wget))
			ids = append(ids, song)
		}
	}
	return ids
}
