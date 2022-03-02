package main

import (
	"go_learn/src/crawler_parallel/engine"
	"go_learn/src/crawler_parallel/scheduler"
	"go_learn/src/crawler_parallel/zhenai/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	e := engine.ConcurrentEngine{Scheduler: &scheduler.SimpleScheduler{}, WorkerCount: 10}
	e.Run(engine.Request{Url: url, ParserFunc: parser.ParseCity})

}
