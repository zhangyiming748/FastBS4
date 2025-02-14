package openkylin

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/soup"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	HOST = "https://gitee.com/organizations/openkylin/projects"
)

func GetAllPages() {
	all := []string{}
	for i := 0; i < 195; i++ {
		log.Printf("正在处理第%d页\n", i+1)
		page := strconv.Itoa(i)
		all = append(all, GetList(page)...)
		time.Sleep(3 * time.Second)
	}
	file, err := os.OpenFile("repos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	for _, page := range all {
		file.WriteString(fmt.Sprintf("%s\n", page))
	}
}
func GetList(page string) (repos []string) {
	host := strings.Join([]string{HOST, page}, "?")
	if page == "0" {
		host = HOST
	}
	html, err := soup.Get(host)
	if err != nil {
		return
	}

	root := soup.HTMLParse(html)
	as := root.FindAll("a")
	for _, a := range as {
		href := a.Attrs()["href"]
		if strings.HasPrefix(href, "/openkylin/") && strings.Count(href, "/") == 2 {
			fmt.Println(href)
			repos = append(repos, href)
		}
	}
	log.Printf("符合过滤条件的仓库名%v\n共%d\n", repos, len(repos))
	return repos
}
