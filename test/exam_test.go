package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

func TestExam(t *testing.T) {
	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	// 确保结束后运行close
	defer resp.Body.Close()

	// gopm get -g -v golang.org/x/text，下载处理乱码的库
	// 处理编码问题，此时是在已知到网页编码为GBK
	//utf8Reader := transform.NewReader(resp.Body,
	//	simplifiedchinese.GBK.NewDecoder())

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,
		e.NewDecoder())

	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			panic(err)
		}
		printCityList(all)
	}
}

// 判断爬取网页内容的编码
func determineEncoding(
	r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	// gopm get -g -v golang.org/x/net，安装库
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}

// 爬取指定的网页信息
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		// m[0]为整个匹配串, matches返回的是 [][][]byte
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}