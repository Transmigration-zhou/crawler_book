package parser

import (
	"crawler_book/engine"
	"crawler_book/model"
	"regexp"
	"strconv"
)

var (
	authorRe    = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
	publisherRe = regexp.MustCompile(`<span class="pl">出版社:</span>[\d\D]*?<a.*?>([^<]+)</a>`)
	pagesRe     = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
	priceRe     = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br/>`)
	scoreRe     = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+) </strong>`)
	introRe     = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)
	idUrlRe     = regexp.MustCompile(`https://book.douban.com/subject/([0-9]+)/`)
)

func ParseBookDetail(contents []byte, name string, url string) engine.ParseResult {
	book := model.Book{}
	book.Name = name
	book.Author = extractContent(contents, authorRe)
	book.Publisher = extractContent(contents, publisherRe)
	if pages, err := strconv.Atoi(extractContent(contents, pagesRe)); err == nil {
		book.Pages = pages
	}
	book.Price = extractContent(contents, priceRe)
	if score, err := strconv.ParseFloat(extractContent(contents, scoreRe), 64); err == nil {
		book.Score = score
	}
	book.Intro = extractContent(contents, introRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "douban",
				Id:      extractContent([]byte(url), idUrlRe),
				Payload: book,
			},
		},
	}
	return result
}

func extractContent(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

type BookDetailParser struct {
	name string
}

func (p *BookDetailParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseBookDetail(contents, url, p.name)

}

func (p *BookDetailParser) Serialize() (name string, args interface{}) {
	return "ParseBookDetail", p.name
}

func NewBookDetailParser(name string) *BookDetailParser {
	return &BookDetailParser{
		name: name,
	}
}
