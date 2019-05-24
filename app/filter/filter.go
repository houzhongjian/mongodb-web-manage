package filter

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Login .
func Login(c *gin.Context) {
	isLogin := sessions.Default(c).Get("isLogin")
	if isLogin == nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	loginStatus := fmt.Sprintf("%s", sessions.Default(c).Get("isLogin"))
	if loginStatus != "yes" {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
}
