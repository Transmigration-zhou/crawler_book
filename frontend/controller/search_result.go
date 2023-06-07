package controller

import (
	"context"
	"crawler_book/engine"
	"crawler_book/frontend/model"
	"crawler_book/frontend/view"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v5"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elasticsearch.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

const pageSize = 10

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q

	resp, err := h.client.Search(
		h.client.Search.WithContext(context.Background()),
		h.client.Search.WithIndex("crawler_book"),
		h.client.Search.WithQuery(rewriteQueryString(q)),
		h.client.Search.WithFrom(from),
	)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return result, errors.New(strconv.Itoa(resp.StatusCode))
	}
	var r map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return result, err
	}
	result.Hits = int64(r["hits"].(map[string]interface{})["total"].(float64))
	result.Start = from
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var item engine.Item
		date, _ := json.Marshal(hit.(map[string]interface{})["_source"])
		json.Unmarshal(date, &item)
		result.Items = append(result.Items, item)
	}
	if result.Start == 0 {
		result.PrevFrom = -1
	} else {
		result.PrevFrom = (result.Start - 1) / pageSize * pageSize
	}
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

// like "Author" to "Payload.Author"
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
