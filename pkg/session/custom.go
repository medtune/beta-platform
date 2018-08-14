package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetLoginStatus from gin.Context
func GetLoginStatus(c *gin.Context) bool {
	if v := sessions.Default(c).Get("logged"); v != nil {
		return v.(bool)
	}
	return false
}

// SetLoginStatus on gin.Context
func SetLoginStatus(c *gin.Context, status bool) {
	s := sessions.Default(c)
	s.Set("logged", status)
	s.Save()
}
