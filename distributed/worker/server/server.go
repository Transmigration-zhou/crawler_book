package main

import (
	"crawler_book/distributed/config"
	"crawler_book/distributed/rpcsupport"
	"crawler_book/distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
