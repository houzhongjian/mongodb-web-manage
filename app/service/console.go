package service

import (
	"log"
	"net/http"

	"github.com/houzhongjian/mongodb-web-manage/app/db"
	"github.com/gin-gonic/gin"
)

//HandleConsole .
func (s *Service) HandleConsole(c *gin.Context) {
	log.Println(db.DB)
	//获取所有的数据库.
	dbList, err := db.Sess.DatabaseNames()
	if err != nil {
		log.Println("数据库获取失败")
		return
	}

	obj := map[string]interface{}{
		"dblist": dbList,
	}
	c.HTML(http.StatusOK, "console.html", obj)
}

//HandleDefault .
func (s Service) HandleDefault(c *gin.Context) {
	c.Redirect(http.StatusFound, "/console")
}
