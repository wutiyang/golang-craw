package engine

import (
	"demoCrawler/fetchr"
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run(seeds ...Request)  {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		ParseResult, err := work(r)
		if err != nil {
			continue
		}

		requests = append(requests, ParseResult.Requests...)

		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func work(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetchr.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error " + "fetching url %s: %v", r.Url, err)

		//continue
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}