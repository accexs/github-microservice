package status

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	alive = "OK"
)

func Alive(c *gin.Context) {
	c.String(http.StatusOK, alive)
}
