package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.you85.net/"
	selector := "li[style='width: 665px; height: 330px;']"

	results, err := scrapeUrl(url, selector)
	if err != nil {
		log.Fatal(err)
	}

	content := strings.Join(results, "\n")
	err = ioutil.WriteFile("output.txt", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("写入文件成功！")
}

func scrapeUrl(url string, selector string) ([]string, error) {
	// 发送HTTP GET请求并获取网页内容
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch webpage: %s", res.Status)
	}

	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	// 提取指定的HTML标签中的内容
	var results []string
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		results = append(results, s.Text())
	})

	return results, nil
}
