package client

import (
	"crawler_book/distributed/config"
	"crawler_book/distributed/rpcsupport"
	"crawler_book/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return out, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Got item#%d %v", itemCount, item)
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("ItemSaver error saving err: %v", err)
			}
		}
	}()
	return out, nil
}
