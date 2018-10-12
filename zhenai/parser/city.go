package parser

import (
	"go-spider/engine"
	"regexp"
)

// 城市解析器
// 得到各个城市首页上的用户名称和用户详情页url
var (
	profileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(
		contents, -1)
	// 对于每个url生成request
	result := engine.ParseResult{}

	for _, m := range matches {
		// 添加内容，如城市名
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		// 添加url的request
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(
				c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(
		contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
