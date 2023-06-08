package persist

import (
	"crawler_book/engine"
	"crawler_book/persist"
	"github.com/elastic/go-elasticsearch/v5"
)

type ItemSaverService struct {
	Client *elasticsearch.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
