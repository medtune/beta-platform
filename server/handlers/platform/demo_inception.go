package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/service/demo"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// InceptionImagenet demo
func InceptionImagenet(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	images, err := demo.CollectImagesData("inception")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	c.Status(200)
	tmpl.DemoInceptionImagenet.Execute(c.Writer, &data.InceptionDemo{
		Main: data.Main{
			Version:   pkg.VERSION,
			PageTitle: "Demo: Image classification",
		},
		Samples: images,
	})
}
