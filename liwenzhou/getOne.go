package liwenzhou

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

// extractArticlePost提取网页中符合article class="post"的内容
func extractArticlePost(root *html.Node) string {
	var articleContent strings.Builder
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "post" {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						html.Render(&articleContent, c)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(root)
	return articleContent.String()
}
func GetOne(uri, title string) {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Printf("请求网页出错: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取网页内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取网页内容出错: %v\n", err)
		return
	}

	// 解析HTML文档
	root, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		fmt.Printf("解析HTML出错: %v\n", err)
		return
	}

	articleText := extractArticlePost(root)
	//log.Printf(articleText)
	// 将提取到的内容写入到markdown文件（这里简单示例，文件名固定为output.md）

	fname := strings.Join([]string{title, "md"}, ".")
	file, err := os.Create(fname)
	if err != nil {
		fmt.Printf("创建文件出错: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(articleText))
	if err != nil {
		fmt.Printf("写入文件出错: %v\n", err)
	}

	fmt.Printf("已提取article class='post'内容并保存为%v文件\n", fname)
}
