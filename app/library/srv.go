package library

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Srv .
type Srv struct {
}

//Jump .
func (s Srv) Jump(path string, c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, path)
}
