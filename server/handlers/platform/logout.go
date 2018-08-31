package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/session"
)

// Logout handler
// always redirect to index
func Logout(c *gin.Context) {
	if logged := session.GetLoginStatus(c); logged {
		session.SetLoginStatus(c, false)
	}
	c.Redirect(302, "/index")
}
