package main

import (
	"go_learn/src/crawler/engine"
	"go_learn/src/crawler/zhenai/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{Url: url, ParserFunc: parser.ParseCity})

}
