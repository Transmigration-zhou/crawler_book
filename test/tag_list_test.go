package test

import (
	"crawler_book/douban/parser"
	"crawler_book/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("https://book.douban.com/")
	if err != nil {
		panic(err)
	}
	result := parser.ParseTagList(contents)

	const resultSize = 47
	expectedUrls := []string{
		"https://book.douban.com/tag/小说",
		"https://book.douban.com/tag/随笔",
		"https://book.douban.com/tag/日本文学",
	}
	expectedTags := []string{
		"小说",
		"随笔",
		"日本文学",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url is %d: %s, but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Items))
	}
	for i, tag := range expectedTags {
		if result.Items[i].(string) != tag {
			t.Errorf("expected tag is %d: %s, but was %s", i, tag, result.Items[i].(string))
		}
	}
}
