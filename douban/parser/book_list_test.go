package parser

import (
	"crawler_book/fetcher"
	"testing"
)

func TestParseBookList(t *testing.T) {
	contents, err := fetcher.Fetch("https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C")
	if err != nil {
		panic(err)
	}
	result := ParseBookList(contents)

	const resultSize = 14
	expectedUrls := []string{
		"https://book.douban.com/subject/36142067/",
		"https://book.douban.com/subject/35407136/",
		"https://book.douban.com/subject/30192800/",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url is %d: %s, but was %s", i, url, result.Requests[i].Url)
		}
	}
}
