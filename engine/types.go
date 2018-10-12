package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 不处理操作的 parser
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
