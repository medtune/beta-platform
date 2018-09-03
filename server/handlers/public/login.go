package public

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Login page handler
func Login(c *gin.Context) {
	// Check session
	// If logged redirect to home
	if session.GetLoginStatus(c) {
		c.Redirect(302, "/home")
		return
	}

	// Render login view
	tmpl.Login.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "login",
	})

}
