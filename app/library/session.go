package library

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Session .
type Session struct {
}

//Set .
func (s Session) Set(key string, val interface{}, c *gin.Context) {
	sess := sessions.Default(c)
	sess.Set(key, val)
	if err := sess.Save(); err != nil {
		log.Println(err)
	}
}

//Get .
func (s Session) Get(key string, c *gin.Context) interface{} {
	sess := sessions.Default(c)
	return sess.Get(sess)
}