package main

import (
	"github.com/houzhongjian/mongodb-web-manage/app/core"
)

func main() {
	engine := core.Engine{}
	engine.RegisterRouter().Run()
}
