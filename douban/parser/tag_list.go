package parser

import (
	"crawler_book/engine"
	"regexp"
)

const tagListRe = `<a href="([^"]+)" class="tag">([^<]+)</a>`

func ParseTagList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(tagListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//fmt.Printf("tag: %s, url: %s \n", m[2], m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://book.douban.com" + string(m[1]),
			ParserFunc: ParseBookList,
		})
	}
	//fmt.Printf("matches found: %d\n", len(matches))
	return result
}
