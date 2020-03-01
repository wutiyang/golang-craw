package main

import (
	"demoCrawler/engine"
	"demoCrawler/persist"
	"demoCrawler/scheduler"
	"demoCrawler/zhenai/parser"
	"fmt"
	"regexp"
)

func main() {
	//engine.SimpleEngine{}.Run(
	//	engine.Request{
	//		Url:"http://www.zhenai.com/zhenghun",
	//		ParserFunc:parser.ParseCityList,
	//	})

	//e := engine.ConcurrentEngine{
	//	Scheduler:&scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//}

	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
	//engine.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun/aba",
	//	ParserFunc:parser.ParseCity,
	//})
	//all, err := fetchr.Fetch("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//
	//parser.ParseCityList(all)

	//printCityList(all)
	//fmt.Printf("%s\n", all)

}

// 转码
// todo

// 获取城市列表
func printCityList(contents []byte)  {
	re := regexp.MustCompile(`a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		fmt.Printf("City:%s, Url:%s \n", m[2], m[1])
	}
}
