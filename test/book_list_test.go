package test

import (
	"crawler_book/douban/parser"
	"crawler_book/fetcher"
	"testing"
)

func TestParseBookList(t *testing.T) {
	contents, err := fetcher.Fetch("https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C")
	if err != nil {
		panic(err)
	}
	result := parser.ParseBookList(contents)

	const resultSize = 14
	expectedUrls := []string{
		"https://book.douban.com/subject/36142067/",
		"https://book.douban.com/subject/35407136/",
		"https://book.douban.com/subject/36237507/",
	}
	expectedBooks := []string{
		"动手学深度学习（PyTorch版）",
		"你看起来好像……我爱你",
		"BERT基础教程：Transformer大模型实战",
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
	for i, book := range expectedBooks {
		if result.Items[i].(string) != book {
			t.Errorf("expected book is %d: %s, but was %s", i, book, result.Items[i].(string))
		}
	}
}
