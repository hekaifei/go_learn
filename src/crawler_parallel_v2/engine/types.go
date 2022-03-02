package engine

import "golang.org/x/net/html"

type Request struct {
	Url        string
	ParserFunc func(*html.Node) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type Profile struct {
	Name      string
	UserUrl   string
	UserId    string
	Age       string
	Gender    string
	Marriage  string
	Location  string
	Height    string
	Education string
	ImageUrl  string
	Introduce string
}

func NilParser(*html.Node) ParseResult {
	return ParseResult{}
}
