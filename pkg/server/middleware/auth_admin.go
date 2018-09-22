package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/session"
)

// ProtectedAdmin .
func ProtectedAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isadmin := session.IsAdmin(c); !isadmin {
			c.Abort()
			c.Redirect(302, "/home")
			return
		}
	}
}
