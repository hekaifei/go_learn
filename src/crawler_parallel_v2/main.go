package main

import (
	"go_learn/src/crawler_parallel_v2/engine"
	"go_learn/src/crawler_parallel_v2/scheduler"
	"go_learn/src/crawler_parallel_v2/zhenai/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	e := engine.ConcurrentEngine{Scheduler: &scheduler.QueueScheduler{}, WorkerCount: 10}
	e.Run(engine.Request{Url: url, ParserFunc: parser.ParseCity})

}
