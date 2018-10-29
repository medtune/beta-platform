package platform

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/service/demo"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// Mura Demo
func Mura(c *gin.Context) {
	c.Status(200)
	tmpl.DemoMura.Execute(c.Writer, &data.Main{
		Version:   internal.VERSION,
		PageTitle: "Demo: MURA Classification",
	})
}

// MuraV2 demo
func MuraV2(c *gin.Context) {
	images, err := demo.CollectImagesData("mura")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}
	c.Status(200)
	tmpl.DemoMuraV2.Execute(c.Writer, &data.MuraV2Demo{
		Main: data.Main{
			Version:   internal.VERSION,
			PageTitle: "Demo V2: MURA Classification",
		},
		Samples: images,
	})
}

// MuraUpload .
func MuraUpload(c *gin.Context) {
	// read file
	file, err := c.FormFile("file")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	ext := filepath.Ext(file.Filename)
	// Verify supported extentions
	if ext != ".png" && ext != ".jpeg" && ext != ".jpg" {
		c.Redirect(302, "/error/401")
		return
	}

	// Use new name if not empty else old filename
	name := c.PostForm("name")
	if name == "" {
		name = file.Filename
	} else {
		name += ext
	}

	savePath := fmt.Sprintf("./static/demos/%s/images/%s", "mura", name)

	// Save file
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	c.Redirect(302, "/demos/mura.v2")
}
