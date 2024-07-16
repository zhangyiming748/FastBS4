package util

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func HttpGetProxy(proxyURL, uri string, outName string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	// 设置代理
	proxy, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}
	client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Failed to retrieve image")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(outName, body, 0644) // 将图片保存为output.jpg
	if err != nil {
		log.Fatalf("Error:%v\n", err)
		return nil, err
	}

	//fmt.Println("Image saved successfully")
	return body, nil
}
