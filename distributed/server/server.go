package main

import (
	"crawler_book/distributed/config"
	"crawler_book/distributed/persist"
	"crawler_book/distributed/rpcsupport"
	"fmt"
	"github.com/elastic/go-elasticsearch/v5"
	"log"
)

func main() {
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
