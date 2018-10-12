package main

import (
	"go-spider/zhenai/parser"
	"io/ioutil"
	"testing"
)

// 测试 CityListParser
func TestParseCityList(t *testing.T) {
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	// 读取本地文件
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := parser.ParseCityList(contents)

	const resultSize = 470
	// 例举几个 url和 city测试
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d",
			resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d",
			resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		// interface{}.(type) => 转换类型
		if result.Items[i].(string) != city {
			t.Errorf("expected url #%d: %s; but "+
				"was %s",
				i, city, result.Items[i].(string))
		}
	}
}
