package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/service/demo"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Mura Demo
func Mura(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	c.Status(200)
	tmpl.DemoMura.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: MURA Classification",
	})
}

// MuraV2 demo
func MuraV2(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	images, err := demo.CollectImagesData("mura")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}
	c.Status(200)
	tmpl.DemoMuraV2.Execute(c.Writer, &data.MuraV2Demo{
		Main: data.Main{
			Version:   pkg.VERSION,
			PageTitle: "Demo V2: MURA Classification",
		},
		Samples: images,
	})
}
