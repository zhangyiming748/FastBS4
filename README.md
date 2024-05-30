# FastBS4
# soup
[![Build Status](https://travis-ci.org/anaskhan96/soup.svg?branch=master)](https://travis-ci.org/anaskhan96/soup)
[![GoDoc](https://godoc.org/github.com/anaskhan96/soup?status.svg)](https://pkg.go.dev/github.com/anaskhan96/soup)
[![Go Report Card](https://goreportcard.com/badge/github.com/anaskhan96/soup)](https://goreportcard.com/report/github.com/anaskhan96/soup)

**Web Scraper in Go, similar to BeautifulSoup**

*soup* is a small web scraper package for Go, with its interface highly similar to that of BeautifulSoup.

Exported variables and functions implemented till now :
```go
var Headers map[string]string // Set headers as a map of key-value pairs, an alternative to calling Header() individually 将 headers 设置为键值对的映射，这是单独调用 Header() 的替代方法
var Cookies map[string]string // Set cookies as a map of key-value  pairs, an alternative to calling Cookie() individually 将 cookie 设置为键值对的映射，这是单独调用 Cookie() 的替代方法
func Get(string) (string,error) {} // Takes the url as an argument, returns HTML string 将 url 作为参数，返回 HTML 字符串
func GetWithClient(string, *http.Client) {} // Takes the url and a custom HTTP client as arguments, returns HTML string 将 url 和自定义 HTTP 客户端作为参数，返回 HTML 字符串
func Post(string, string, interface{}) (string, error) {} // Takes the url, bodyType, and payload as an argument, returns HTML string 将 url、bodyType 和 Payload 作为参数，返回 HTML 字符串
func PostForm(string, url.Values) {} // Takes the url and body. bodyType is set to "application/x-www-form-urlencoded" 获取 url 和正文。 bodyType 设置为“application/x-www-form-urlencoded”
func Header(string, string) {} // Takes key,value pair to set as headers for the HTTP request made in Get() 将键值对设置为 Get() 中发出的 HTTP 请求的标头
func Cookie(string, string) {} // Takes key, value pair to set as cookies to be sent with the HTTP request in Get() 在 Get() 中将键、值对设置为与 HTTP 请求一起发送的 cookie
func HTMLParse(string) Root {} // Takes the HTML string as an argument, returns a pointer to the DOM constructed 将 HTML 字符串作为参数，返回指向构造的 DOM 的指针
func Find([]string) Root {} // Element tag,(attribute key-value pair) as argument, pointer to first occurence returned 元素标签，（属性键值对）作为参数，返回指向第一次出现的指针
func FindAll([]string) []Root {} // Same as Find(), but pointers to all occurrences returned 与 Find() 相同，但返回指向所有出现次数的指针
func FindStrict([]string) Root {} //  Element tag,(attribute key-value pair) as argument, pointer to first occurence returned with exact matching values 元素标签，（属性键值对）作为参数，指向第一次出现的指针，返回精确匹配的值
func FindAllStrict([]string) []Root {} // Same as FindStrict(), but pointers to all occurrences returned 与 FindStrict() 相同，但返回指向所有出现次数的指针
func FindNextSibling() Root {} // Pointer to the next sibling of the Element in the DOM returned 指向 DOM 中返回的元素的下一个同级的指针
func FindNextElementSibling() Root {} // Pointer to the next element sibling of the Element in the DOM returned 指向返回的 DOM 中 Element 的下一个同级元素的指针
func FindPrevSibling() Root {} // Pointer to the previous sibling of the Element in the DOM returned 指向 DOM 中返回的元素的前一个同级的指针
func FindPrevElementSibling() Root {} // Pointer to the previous element sibling of the Element in the DOM returned 指向 DOM 中返回的 Element 的上一个同级元素的指针
func Children() []Root {} // Find all direct children of this DOM element 查找此 DOM 元素的所有直接子元素
func Attrs() map[string]string {} // Map returned with all the attributes of the Element as lookup to their respective values 返回的映射包含元素的所有属性，作为对其各自值的查找
func Text() string {} // Full text inside a non-nested tag returned, first half returned in a nested one 返回非嵌套标签内的全文，前半部分在嵌套标签中返回
func FullText() string {} // Full text inside a nested/non-nested tag returned 返回嵌套/非嵌套标签内的全文
func SetDebug(bool) {} // Sets the debug mode to true or false; false by default 将调试模式设置为 true 或 false；默认为 false
func HTML() {} // HTML returns the HTML code for the specific element //HTML 返回特定元素的 HTML 代码
```

`Root` is a struct, containing three fields :
* `Pointer` containing the pointer to the current html node
* `NodeValue` containing the current html node's value, i.e. the tag name for an ElementNode, or the text in case of a TextNode
* `Error` containing an error in a struct if one occurrs, else `nil` is returned.
  A detailed text explaination of the error can be accessed using the `Error()` function. A field `Type` in this struct of type `ErrorType` will denote the kind of error that took place, which will consist of either of the following
    * `ErrUnableToParse`
    * `ErrElementNotFound`
    * `ErrNoNextSibling`
    * `ErrNoPreviousSibling`
    * `ErrNoNextElementSibling`
    * `ErrNoPreviousElementSibling`
    * `ErrCreatingGetRequest`
    * `ErrInGetRequest`
    * `ErrReadingResponse`

## Installation
Install the package using the command
```bash
go get github.com/anaskhan96/soup
```

## Example
An example code is given below to scrape the "Comics I Enjoy" part (text and its links) from [xkcd](https://xkcd.com).

[More Examples](https://github.com/anaskhan96/soup/tree/master/examples)
```go
package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {
	resp, err := soup.Get("https://xkcd.com")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "comicLinks").FindAll("a")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}
```

## Contributions
This package was developed in my free time. However, contributions from everybody in the community are welcome, to make it a better web scraper. If you think there should be a particular feature or function included in the package, feel free to open up a new issue or pull request.