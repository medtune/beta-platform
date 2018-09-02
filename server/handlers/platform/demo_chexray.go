package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/service/demo"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Chexray demo
func Chexray(c *gin.Context) {
	c.Status(200)
	tmpl.DemoChexray.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Demo: Chest X-Ray Classification",
	})
}

// ChexrayV2 demo
func ChexrayV2(c *gin.Context) {
	images, err := demo.CollectImagesData("chexray")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	c.Status(200)
	tmpl.DemoChexrayV2.Execute(c.Writer, &data.ChexrayV2Demo{
		Main: data.Main{
			Version:   pkg.VERSION,
			PageTitle: "Demo V2: Chest X-Ray Classification",
		},
		Samples: images,
	})
}
