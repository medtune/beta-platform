package public

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/session"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
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
