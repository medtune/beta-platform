package public

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/session"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// Index page handler
func Index(c *gin.Context) {
	if session.GetLoginStatus(c) {
		c.Redirect(302, "home")
	} else {
		c.Status(200)
		inject := data.Main{
			Version:   internal.VERSION,
			PageTitle: "Index",
		}
		tmpl.Index.Execute(c.Writer, &inject)
	}
}
