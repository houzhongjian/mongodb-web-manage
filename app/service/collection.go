package service

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/houzhongjian/mongodb-web-manage/app/db"
	"github.com/gin-gonic/gin"
)

//HandleCollections .
func (s Service) HandleCollections(c *gin.Context) {
	log.Println(123)
	dbname := c.Param("database")
	dbList, err := db.Sess.DatabaseNames()
	if err != nil {
		log.Println("数据库获取失败")
		return
	}

	collList, err := db.Sess.DB(dbname).CollectionNames()
	if err != nil {
		log.Println("数据库获取失败")
		return
	}

	obj := map[string]interface{}{
		"dbname":   dbname,
		"dblist":   dbList,
		"collList": collList,
	}
	c.HTML(http.StatusOK, "collection.html", obj)
}

//HandleCommand .
func (s Service) HandleCommand(c *gin.Context) {
	dbname := c.Param("database")
	db.DB.Database = dbname
	collname := c.Param("collection")

	dbList, err := db.Sess.DatabaseNames()
	if err != nil {
		log.Println("数据库获取失败")
		return
	}

	collList, err := db.Sess.DB(dbname).CollectionNames()
	if err != nil {
		log.Println("数据库获取失败")
		return
	}

	shell := c.Query("shell")

	shell, _ = url.PathUnescape(shell)
	if len(shell) < 1 {
		shell = fmt.Sprintf("db.%s.find()", collname)
	}

	shell = s.Mgo.ParseShell(shell)
	log.Println("run shell:", shell)
	js := s.Mgo.Query(shell)

	obj := map[string]interface{}{
		"shell":    shell,
		"dbname":   dbname,
		"collname": collname,
		"dblist":   dbList,
		"collList": collList,
		"js":       js,
	}
	c.HTML(http.StatusOK, "command.html", obj)
}
