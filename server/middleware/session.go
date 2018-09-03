package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/session"
)

// Session .
func Session() gin.HandlerFunc {
	return session.DefaultHandler()
}
