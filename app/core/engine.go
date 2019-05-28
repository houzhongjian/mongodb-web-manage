package core

import (
	"github.com/houzhongjian/mongodb-web-manage/app/service"
	"github.com/gin-gonic/gin"
)

//Engine .
type Engine struct {
	Route *Router
}

//Router .
type Router struct {
	Router *gin.Engine
	Srv    service.Service
}

//NewEngine .
func NewEngine() Engine {
	return Engine{}
}
