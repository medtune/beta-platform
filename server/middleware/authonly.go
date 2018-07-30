package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//FIXME
func LoggedOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		if v := s.Get("logged"); v == nil || !v.(bool) {
			c.AbortWithStatus(404)
			return
		}
	}
}
