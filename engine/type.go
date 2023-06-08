package engine

type Parser interface {
	Parse([]byte, string) ParseResult
	Serialize() (string, interface{})
}

type ParserFunc func([]byte, string) ParseResult

type Request struct {
	Url string
	//ParserFunc ParserFunc
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct{}

func (*NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (*NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (string, interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
