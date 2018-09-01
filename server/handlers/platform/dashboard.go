package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Dashboard .
func Dashboard(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}

	c.Status(200)
	tmpl.Dashboard.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Dashboard",
	})
}
