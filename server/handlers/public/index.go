package public

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func Index(c *gin.Context) {
	if session.GetLoginStatus(c) {
		c.Redirect(302, "home")
	} else {
		c.Status(200)
		inject := data.Main{
			Version:   pkg.VERSION,
			PageTitle: "Index",
		}
		tmpl.Index.Execute(c.Writer, &inject)
	}
}
