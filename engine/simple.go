package engine

import (
	"go-spider/fetcher"
	"log"
)

type SimpleEngine struct {}

// 执行，根据种子 request
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 循环 request爬取解析
	for len(requests) > 0 {
		r := requests[0]
		// 去掉 requests第一个值
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

// 分布式爬虫
func worker(r Request) (ParseResult, error)  {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}