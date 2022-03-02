package engine

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"go_learn/src/crawler_parallel_v2/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher： error，fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	node, _ := htmlquery.Parse(bytes.NewReader(body))
	return r.ParserFunc(node), nil
}
