package parser

import (
	"go-spider/engine"
	"regexp"
)

// 城市列表解析器
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	// 对于每个url生成request
	result := engine.ParseResult{}

	for _, m := range matches {
		// 添加内容，如城市名
		result.Items = append(result.Items, "City "+string(m[2]))
		// 添加url的request
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
