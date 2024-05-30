package rss

import (
	"fmt"
	"github.com/zhangyiming748/FastBS4/util"
	"os/exec"
	"strings"
)

type one struct {
	name string // 文件名
	uri  string // 地址
}

func Rss() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	file := []string{}
	lines := util.ReadByLine("rss.sh")
	for _, line := range lines {
		if strings.Contains(line, "<title>") {
			fmt.Println(line)
			title := strings.Split(line, "《")[1]
			fmt.Println("1", title)
			title = strings.Split(title, "》")[0]
			fmt.Println("2", title)
			title = strings.Join([]string{title, "mp3"}, ".")
			fmt.Println("3", title)
			file = append(file, title)
		}
		if strings.Contains(line, "<enclosure url=") {
			fmt.Println(line)
			prefix := strings.Split(line, "\" length")[0]
			fmt.Println("4", prefix)
			suffix := strings.Split(prefix, "url=\"")[1]
			fmt.Println("5", suffix)
			file = append(file, suffix)
		}
	}
	//util.WriteByLine("rss.txt", file)
}
func download() {
	cmds := []string{}
	f := util.ReadByLine("rss.txt")
	for i := 0; i < len(f)-1; i += 2 {
		name := f[i]
		uri := f[i+1]
		fmt.Printf("文件名:%s\t地址%s\n", name, uri)
		//alias wget-proxy="wget -e use_proxy=yes -e http_proxy=192.168.1.10:8889 -e https_proxy=192.168.1.10:8889"
		cmd := exec.Command("wget", "-e", "use_proxy=yes", "-e", "http_proxy=192.168.1.20:8889", "-e", " https_proxy=192.168.1.20:8889", "--no-check-certificate", "-O", name, uri)
		cmds = append(cmds, cmd.String())
	}

	util.WriteByLine("down.sh", cmds)
}
