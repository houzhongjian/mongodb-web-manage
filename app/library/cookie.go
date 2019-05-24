package library

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Cookie struct {
}

//CloseAll .
func (s Cookie) CloseAll(c *gin.Context) {
	cookieList := c.Request.Cookies()
	for _, v := range cookieList {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    v.Name,
			Expires: time.Now().AddDate(-1, 0, 0),
		})
	}
}
