package engine

import (
	"bytes"
	"go_learn/src/crawler/fetcher"
	"log"
	"time"

	"github.com/antchfx/htmlquery"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	num := 0

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher： error，fetching url %s: %v", r.Url, err)
			continue
		}
		node, _ := htmlquery.Parse(bytes.NewReader(body))
		parseResult := r.ParserFunc(node)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			num++
			log.Printf("item %d: %s", num, item)
		}
		time.Sleep(1000)
	}

}
