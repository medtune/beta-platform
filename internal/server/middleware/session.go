package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/session"
)

// Session .
func Session() gin.HandlerFunc {
	return session.DefaultHandler()
}
