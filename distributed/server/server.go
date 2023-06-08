package main

import (
	"crawler_book/distributed/config"
	"crawler_book/distributed/persist"
	"crawler_book/distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v5"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ESIndex))
}

func serveRpc(host, index string) error {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
