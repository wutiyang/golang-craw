package parser

import (
	"demoCrawler/engine"
	"fmt"
	"regexp"
)

var (
	cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	submatchs := cityRe.FindAllSubmatch(contents, -1)

	fmt.Printf("%s", submatchs)
	result := engine.ParseResult{}
	for _,match := range submatchs {
		name := string(match[2])
		result.Items = append(result.Items, "User" + name)

		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	// 分页数据
	matches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseCity,
		})
	}

	return result
}
