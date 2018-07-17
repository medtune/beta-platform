package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/session"
)

func Session() gin.HandlerFunc {
	return session.DefaultHandler()
}
