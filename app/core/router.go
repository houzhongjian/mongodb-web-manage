package core

import (
	"fmt"
	"time"

	"dbclient.net/web/mongodb/app/filter"
	"dbclient.net/web/mongodb/app/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//RegisterRouter .
func (e *Engine) RegisterRouter() *Router {
	e.Route = &Router{
		Router: gin.Default(),
		Srv:    service.Service{},
	}

	e.Route.register()
	return e.Route
}

//Run .
func (r *Router) Run(addr ...string) error {
	port := "9090"
	if len(addr) > 0 {
		port = addr[0]
	}
	return r.Router.Run(fmt.Sprintf(":%s", port))
}

//register .
func (r *Router) register() {
	r.Router.Static("/css/", "./app/public/css/")
	r.Router.Static("/js/", "./app/public/js/")
	r.Router.Static("/img/", "./app/public/img/")
	r.Router.Static("/plugins/", "./app/public/plugins/")

	//=>load tempate
	r.Router.LoadHTMLGlob("./app/view/*")

	//=>session .
	store := cookie.NewStore([]byte("secret"))
	r.Router.Use(sessions.Sessions(fmt.Sprintf("%d", time.Now().UnixNano()), store))

	//=>login.
	r.Router.GET("/login", r.Srv.HandleLogin)
	r.Router.POST("/login", r.Srv.HandleLogin)

	r.Router.Use(filter.Login)
	{
		//=> console
		r.Router.GET("/", r.Srv.HandleDefault)
		r.Router.GET("/console", r.Srv.HandleConsole)

		//=> collections
		r.Router.GET("/database/:database", r.Srv.HandleCollections)
		r.Router.GET("/database/:database/collection/:collection", r.Srv.HandleCommand)
	}
}
