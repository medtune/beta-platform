package middleware

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/jsonutil"
)

// ProtectedView .
func ProtectedView() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		if v := s.Get("logged"); v == nil || !v.(bool) {
			c.Abort()
			c.Redirect(302, "/index")
			return
		}
	}
}

// ProtectedAPI .
func ProtectedAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		//FIXME
		s := sessions.Default(c)
		if v := s.Get("logged"); v == nil || !v.(bool) {
			c.AbortWithStatusJSON(200, jsonutil.Fail(
				fmt.Errorf("protected api ep <authentification required>")),
			)
			return
		}
	}
}
