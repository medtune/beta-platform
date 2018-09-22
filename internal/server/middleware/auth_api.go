package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/beta-platform/internal/session"
)

// ProtectedAPI .
func ProtectedAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		if v := session.GetLoginStatus(c); !v {
			c.AbortWithStatusJSON(200, jsonutil.Fail(
				fmt.Errorf("protected api ep: <authentification required>")),
			)
			return
		}
	}
}
