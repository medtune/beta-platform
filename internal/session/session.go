package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	SessID string         = "default-sessions"
	Store  sessions.Store = nil
)

// NewStore .
func NewStore(secret string) sessions.Store {
	return cookie.NewStore([]byte(secret))
}

// NewHandler .
func NewHandler(name string, store sessions.Store) gin.HandlerFunc {
	return sessions.Sessions(name, store)
}

// DefaultHandler .
func DefaultHandler() gin.HandlerFunc {
	return NewHandler(SessID, Store)
}
