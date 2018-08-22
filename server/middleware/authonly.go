package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoggedOnly .
func LoggedOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		//FIXME
		s := sessions.Default(c)
		if v := s.Get("logged"); v == nil || !v.(bool) {
			c.Abort()
			c.Redirect(302, "/index")
			return
		}
	}
}
