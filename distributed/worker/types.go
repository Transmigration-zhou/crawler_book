package worker

import (
	"crawler_book/distributed/config"
	"crawler_book/douban/parser"
	"crawler_book/engine"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser1, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser1,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserilizing request:%v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}

	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseTagList:
		return engine.NewFuncParser(
			parser.ParseTagList,
			config.ParseTagList), nil
	case config.ParseBookList:
		return engine.NewFuncParser(
			parser.ParseBookList,
			config.ParseBookList), nil
	case config.NilParser:
		return &engine.NilParser{}, nil
	case config.ParseBookDetail:
		if name, ok := p.Args.(string); ok {
			return parser.NewBookDetailParser(name), nil
		} else {
			return nil, fmt.Errorf("invalid arg:%v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}
