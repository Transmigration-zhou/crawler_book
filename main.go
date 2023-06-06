package main

import (
	"crawler_book/douban/parser"
	"crawler_book/engine"
	"crawler_book/persist"
	"crawler_book/scheduler"
)

func main() {
	//simpleEngine := engine.SimpleEngine{}
	//simpleEngine.Run(engine.Request{
	//	Url:        "https://book.douban.com/",
	//	ParserFunc: parser.ParseTagList,
	//})

	saver, err := persist.ItemSaver("crawler_book")
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 5,
		ItemChan:  saver,
	}
	concurrentEngine.Run(engine.Request{
		//Url:        "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
		//ParserFunc: parser.ParseBookList,
		Url:        "https://book.douban.com/",
		ParserFunc: parser.ParseTagList,
	})
}
