package hidden

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func Home(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	inject := data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Home",
	}
	c.Status(200)
	tmpl.Home.Execute(c.Writer, &inject)
}
