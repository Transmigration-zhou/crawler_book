package main

import (
	"crawler_book/distributed/client"
	"crawler_book/distributed/config"
	"crawler_book/distributed/rpcsupport"
	client2 "crawler_book/distributed/worker/client"
	"crawler_book/douban/parser"
	"crawler_book/engine"
	"crawler_book/scheduler"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "", "worker hosts (comma separated)") //逗号分割
)

func main() {
	flag.Parse()
	saver, err := client.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := client2.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkCount:        5,
		ItemChan:         saver,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		//Url: "https://book.douban.com/tag/%E9%80%9A%E4%BF%A1",
		//Parser: engine.NewFuncParser(
		//	parser.ParseBookList,
		//	config.ParseBookList,
		//),
		Url: "https://book.douban.com/",
		Parser: engine.NewFuncParser(
			parser.ParseBookList,
			config.ParseBookList,
		),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
