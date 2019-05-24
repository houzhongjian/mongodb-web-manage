package main

import (
	"dbclient.net/web/mongodb/app/core"
)

func main() {
	engine := core.Engine{}
	engine.RegisterRouter().Run()
}
