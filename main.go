package main

import (
	"crawler_book/douban/parser"
	"crawler_book/engine"
)

func main() {
	simpleEngine := engine.SimpleEngine{}
	simpleEngine.Run(engine.Request{
		Url:        "https://book.douban.com/",
		ParserFunc: parser.ParseTagList,
	})
}
