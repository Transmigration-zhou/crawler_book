package parser

import (
	"crawler_book/engine"
	"regexp"
)

const bookListRe = `<a href="([^"]+)" title="([^"]+)"`

func ParseBookList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(bookListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		//fmt.Printf("book: %s, url: %s \n", m[2], m[1])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(content []byte) engine.ParseResult {
				return ParseBookDetail(content, name)
			},
		})
	}
	//fmt.Printf("matches found: %d\n", len(matches))
	return result
}
