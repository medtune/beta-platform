package public

import (
	"github.com/gin-gonic/gin"
)

func NoRouteProxy(c *gin.Context) {
	c.Redirect(302, "/error/404")
}
