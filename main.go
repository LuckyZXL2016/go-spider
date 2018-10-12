package main

import (
	"go-spider/engine"
	"go-spider/scheduler"
	"go-spider/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	// 爬取所有城市
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// 爬取单一城市，如上海
	//e.Run(engine.Request{
	//		Url:    "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
