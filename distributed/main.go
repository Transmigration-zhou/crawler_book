package main

import (
	"crawler_book/distributed/client"
	"crawler_book/distributed/config"
	client2 "crawler_book/distributed/worker/client"
	"crawler_book/douban/parser"
	"crawler_book/engine"
	"crawler_book/scheduler"
	"fmt"
)

func main() {
	saver, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := client2.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkCount:        5,
		ItemChan:         saver,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "https://book.douban.com/tag/%E9%80%9A%E4%BF%A1",
		Parser: engine.NewFuncParser(
			parser.ParseBookList,
			config.ParseBookList,
		),
		//Url: "https://book.douban.com/",
		//Parser: engine.NewFuncParser(
		//	parser.ParseBookList,
		//	config.ParseBookList,
		//),
	})
}
