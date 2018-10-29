package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/session"
)

// Logout handler
// always redirect to index
func Logout(c *gin.Context) {
	session.Disconnect(c)
	c.Redirect(302, "/index")
}
