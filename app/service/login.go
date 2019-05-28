package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/houzhongjian/mongodb-web-manage/app/db"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

//HandleLogin .
func (s *Service) HandleLogin(c *gin.Context) {
	if c.Request.Method == "POST" {
		host := c.PostForm("host")
		user := c.PostForm("user")
		password := c.PostForm("password")
		port := c.PostForm("port")
		mechanism := c.PostForm("mechanism")

		mongoHost := fmt.Sprintf("%s:%s", host, port)
		session, err := mgo.Dial(mongoHost)
		if err != nil {
			log.Printf("%+v\n", err)
			c.Writer.WriteString("登录失败")
			return
		}

		auth := mgo.Credential{
			Username:  user,
			Password:  password,
			Mechanism: mechanism,
		}
		if err := session.Login(&auth); err != nil {
			log.Printf("%+v\n", err)
			c.Writer.WriteString("认证失败")
			return
		}

		db.Sess = session

		db.DB = &db.DBConfig{
			Host:     host,
			Password: password,
			User:     user,
			Port:     port,
			Auth: db.Auth{
				Mechanism: mechanism,
				Db:        "admin",
			},
		}

		s.Base.Session.Set("host", host, c)
		s.Base.Session.Set("user", user, c)
		s.Base.Session.Set("password", password, c)
		s.Base.Session.Set("isLogin", "yes", c)

		s.Srv.Jump("/console", c)
		return
	}
	//reflash all cookie.
	s.Base.Cookie.CloseAll(c)
	c.HTML(http.StatusOK, "login.html", nil)
}
