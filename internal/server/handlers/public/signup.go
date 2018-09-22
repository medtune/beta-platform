package public

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// Signup page handler
func Signup(c *gin.Context) {
	// Render signup page
	tmpl.Signup.Execute(c.Writer, &data.Main{
		PageTitle: "Signup",
	})
}
