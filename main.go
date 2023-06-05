package main

import (
	"crawler_book/douban/parser"
	"crawler_book/engine"
	"crawler_book/scheduler"
)

func main() {
	//simpleEngine := engine.SimpleEngine{}
	//simpleEngine.Run(engine.Request{
	//	Url:        "https://book.douban.com/",
	//	ParserFunc: parser.ParseTagList,
	//})

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		//Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 5,
	}
	concurrentEngine.Run(engine.Request{
		Url:        "https://book.douban.com/",
		ParserFunc: parser.ParseTagList,
	})
}
