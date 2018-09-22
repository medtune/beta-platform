package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/service/demo"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// InceptionImagenet demo
func InceptionImagenet(c *gin.Context) {
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
