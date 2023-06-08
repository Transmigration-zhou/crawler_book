package main

import (
	"crawler_book/distributed/config"
	"crawler_book/douban/parser"
	"crawler_book/engine"
	"crawler_book/persist"
	"crawler_book/scheduler"
)

func main() {
	saver, err := persist.ItemSaver("crawler_book")
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkCount:        5,
		ItemChan:         saver,
		RequestProcessor: engine.Worker,
	}
	concurrentEngine.Run(engine.Request{
		Url: "https://book.douban.com/",
		Parser: engine.NewFuncParser(
			parser.ParseBookList,
			config.ParseBookList,
		),
	})
}
