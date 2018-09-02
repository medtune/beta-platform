package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/session"
)

// ProtectedView .
func ProtectedView() gin.HandlerFunc {
	return func(c *gin.Context) {
		if v := session.GetLoginStatus(c); !v {
			c.Abort()
			c.Redirect(302, "/login")
			return
		}
	}
}
