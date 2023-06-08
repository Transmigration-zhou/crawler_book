package persist

import (
	"bytes"
	"context"
	"crawler_book/engine"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v5"
	"github.com/elastic/go-elasticsearch/v5/esapi"
	"log"
	"strconv"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return out, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Got item#%d %v", itemCount, item)
			err := Save(client, index, item)
			if err != nil {
				log.Printf("ItemSaver error saving err: %v", err)
			}
		}
	}()
	return out, nil
}

func Save(client *elasticsearch.Client, index string, item engine.Item) error {
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", data)
	indexService := esapi.IndexRequest{
		Index:        index,
		DocumentType: item.Type,
		Body:         bytes.NewReader(data),
	}
	if item.Id != "" {
		indexService.DocumentID = item.Id
	}
	resp, err := indexService.Do(context.Background(), client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return errors.New(strconv.Itoa(resp.StatusCode))
	}
	log.Printf("%+v\n", resp)
	return nil
}
