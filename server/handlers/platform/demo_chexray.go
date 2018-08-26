package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Chexray demo
func Chexray(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	c.Status(200)
	tmpl.DemoChexray.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Chest X-Ray Classification",
	})
}

// ChexrayV2 demo
func ChexrayV2(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	c.Status(200)
	tmpl.DemoChexrayV2.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo V2: Chest X-Ray Classification",
	})
}
